package runtime

func ForEachG() []*G {
	var all []*G
	forEachG(func(gp *g) {
		g := castg(gp)
		all = append(all, g)
	})
	return all
}

func ForEachM() []*M {
	var all []*M
	lock(&sched.lock)
	defer unlock(&sched.lock)
	for mp := allm; mp != nil; mp = mp.alllink {
		m := castm(mp)
		all = append(all, m)
	}
	return all
}

func ForEachP() []*P {
	all := make([]*P, 0, gomaxprocs)
	lock(&sched.lock)
	defer unlock(&sched.lock)
	for _, p := range allp {
		pp := castp(p)
		all = append(all, pp)
	}
	return all
}

func Sched() *Schedt {
	lock(&sched.lock)
	defer unlock(&sched.lock)
	return castsched()
}

func Newproc(fn *funcval, annotations ...string) {
	gp := getg()
	pc := sys.GetCallerPC()
	systemstack(func() {
		newg := newproc1(fn, gp, pc, false, waitReasonZero)
		newg.annotations = annotations

		pp := getg().m.p.ptr()
		runqput(pp, newg, true)

		if mainStarted {
			wakep()
		}
	})
}
