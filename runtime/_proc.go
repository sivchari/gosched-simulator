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
	var all []*P
	lock(&allpLock)
	defer unlock(&allpLock)
	for _, p := range allp {
		gs := make([]G, 0)
		for !runqempty(p) {
			gp, _ := runqget(p)
			g := castg(gp)
			gs = append(gs, *g)
		}
		pp := castp(p)
		pp.XRunq = gs
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
