// Симулятор рантайма Go.
package main

import (
	"fmt"
	"slices"
)

// максимальное время непрерывного выполнения горутины на потоке
const maxRunDur = 100

// статус горутины
type gStatus string

// статусы горутин
var (
	statusRunnable gStatus = "runnable" // готова к выполнению
	statusRunning  gStatus = "running"  // выполняется на потоке
	statusWaiting  gStatus = "waiting"  // заблокирована
	statusDead     gStatus = "dead"     // завершилась
)

// Goroutine представляет горутину.
type Goroutine struct {
	id     int     // идентификатор, нумерация с 1
	runDur int     // время в состоянии running
	status gStatus // статус
}

// Block переводит горутину в состояние waiting.
func (g *Goroutine) Block() {
	if g.status != statusRunning {
		panic("invalid status for block: " + g.status)
	}
	g.status = statusWaiting
}

// Unblock переводит горутину в состояние runnable.
func (g *Goroutine) Unblock() {
	if g.status != statusWaiting {
		panic("invalid status for unblock: " + g.status)
	}
	g.status = statusRunnable
}

// Done переводит горутину в состояние dead.
func (g *Goroutine) Done() {
	if g.status != statusRunning {
		panic("invalid status for done: " + g.status)
	}
	g.status = statusDead
}

// Thread представляет поток операционной системы.
type Thread struct {
	id   int        // идентификатор, нумерация с 1
	goro *Goroutine // горутина на выполнении
}

// RuntimeState представляет состояние рантайма.
type RuntimeState struct {
	dur      int         // общее время выполнения
	threads  map[int]int // ключ - id потока, значение - id горутины (0 - поток свободен)
	runnable []int       // id горутин в очереди на выполнение
	running  []int       // id горутин на выполнении
	waiting  []int       // id заблокированных горутин
	dead     []int       // id завершенных горутин
}

// Runtime представляет симулятор рантайма.
type Runtime struct {
	dur      int          // общее время выполнения
	nGoro    int          // счетчик горутин
	threads  []*Thread    // потоки
	runnable []*Goroutine // горутины в очереди на выполнение
	running  []*Goroutine // горутины на выполнении
	waiting  []*Goroutine // заблокированные горутины
	dead     []*Goroutine // завершенные горутины
}

// NewRuntime создает новый рантайм на gomaxprocs потоках.
func NewRuntime(gomaxprocs int) *Runtime {
	threads := make([]*Thread, gomaxprocs)
	for i := range gomaxprocs {
		threads[i] = &Thread{id: i + 1}
	}
	return &Runtime{threads: threads}
}

// Go создает новую горутину в рантайме.
func (r *Runtime) Go() *Goroutine {
	r.nGoro++
	g := &Goroutine{id: r.nGoro, status: statusRunnable}
	r.runnable = append(r.runnable, g)
	return g
}

// начало решения

// Forward двигает время вперед на dur единиц.
func (r *Runtime) Forward(dur int) {
	r.dur += dur
	for _, goro := range r.running {
		goro.runDur += dur
	}
}

func (r *Runtime) Schedule() {
	var massDelete, idsGoro []int
	for indexGo, goro := range r.running {
		isDelete := false
		// Проверка условия, что горутина выполняется слишком долго
		if goro.runDur > maxRunDur {
			goro.status = statusRunnable
			goro.runDur = 0
			r.runnable = append(r.runnable, goro)
			massDelete = append(massDelete, indexGo)
			idsGoro = append(idsGoro, goro.id)
			isDelete = true
		}
		// если горутина заблокировалась
		if goro.status == statusWaiting {
			r.waiting = append(r.waiting, goro)
			if !isDelete {
				massDelete = append(massDelete, indexGo)
				idsGoro = append(idsGoro, goro.id)
				isDelete = true
			}
		}

		// если горутина завершилась
		if goro.status == statusDead {
			r.dead = append(r.dead, goro)
			if !isDelete {
				massDelete = append(massDelete, indexGo)
				idsGoro = append(idsGoro, goro.id)
				isDelete = true
			}
		}
	}
	if len(massDelete) > 0 {
		// Удалить горутины из выполняемых
		r.running = slices.Delete(r.running, massDelete[0], massDelete[len(massDelete)-1]+1)
		// Освободить потоки
		for indexThread, thread := range r.threads {
			for _, goroId := range idsGoro {
				if thread.goro != nil && thread.goro.id == goroId {
					r.threads[indexThread].goro = nil
				}
			}
		}
		massDelete = []int{}
	}

	for indexGo, goro := range r.waiting {
		if goro.status == statusRunnable {
			r.runnable = append(r.runnable, goro)
			massDelete = append(massDelete, indexGo)
		}
	}
	if len(massDelete) > 0 {
		// Удалить горутины из ожидаемых
		r.waiting = slices.Delete(r.waiting, massDelete[0], massDelete[len(massDelete)-1]+1)
		// Освободить потоки
		for indexThread, thread := range r.threads {
			for _, goroId := range idsGoro {
				if thread.goro != nil && thread.goro.id == goroId {
					r.threads[indexThread].goro = nil
				}
			}
		}
	}

	for indexThread, thread := range r.threads {
		if thread.goro == nil && len(r.runnable) > 0 {
			goro := r.runnable[0]
			goro.status = statusRunning
			r.threads[indexThread].goro = goro
			r.running = append(r.running, goro)
			r.runnable = slices.Delete(r.runnable, 0, 1)
		}
	}
}

// конец решения

// State возвращает текущее состояние рантайма.
func (r *Runtime) State() RuntimeState {
	threads := make(map[int]int)
	for _, t := range r.threads {
		if t.goro != nil {
			threads[t.id] = t.goro.id
		} else {
			threads[t.id] = 0
		}
	}
	runnable := make([]int, len(r.runnable))
	for i, g := range r.runnable {
		runnable[i] = g.id
	}
	running := make([]int, len(r.running))
	for i, g := range r.running {
		running[i] = g.id
	}
	waiting := make([]int, len(r.waiting))
	for i, g := range r.waiting {
		waiting[i] = g.id
	}
	dead := make([]int, len(r.dead))
	for i, g := range r.dead {
		dead[i] = g.id
	}
	return RuntimeState{
		dur:      r.dur,
		threads:  threads,
		runnable: runnable,
		running:  running,
		waiting:  waiting,
		dead:     dead,
	}
}

func main() {

	// создаем рантайм на 2 потока
	r := NewRuntime(2)

	// r.Go()
	g1 := r.Go()
	g2 := r.Go()
	r.Go()
	r.Go()
	r.Schedule()
	
	r.Forward(10)
	g1.Block()
	g2.Block()
	r.Schedule()

	r.Forward(10)
	g1.Unblock()
	g2.Unblock()
	r.Schedule()


	// // создаем 4 горутины
	// g1 := r.Go()
	// g2 := r.Go()
	// r.Go()
	// r.Go()
	// r.Schedule()

	// // прошло 10 единиц времени, g1 завершила выполнение, g2 заблокирована
	// r.Forward(10)
	// g1.Done()
	// g2.Block()
	// r.Schedule()

	// выводим текущее состояние рантайма
	state := r.State()
	fmt.Printf("%+v\n", state)
	// // {dur:10 threads:map[1:3 2:4] runnable:[] running:[3 4] waiting:[2] dead:[1]}
}
