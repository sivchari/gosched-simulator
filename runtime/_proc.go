package runtime

func ForEachG() []*G {
	var all []*G
	forEachG(func(gp *g) {
		g := &G{
			Stack:             gp.stack,
			Stackguard0:       gp.stackguard0,
			Stackguard1:       gp.stackguard1,
			Panic:             gp._panic,
			Defer:             gp._defer,
			M:                 gp.m,
			Sched:             gp.sched,
			Syscallsp:         gp.syscallsp,
			Syscallpc:         gp.syscallpc,
			Syscallbp:         gp.syscallbp,
			Stktopsp:          gp.stktopsp,
			Param:             gp.param,
			Atomicstatus:      gp.atomicstatus,
			StackLock:         gp.stackLock,
			Goid:              gp.goid,
			Schedlink:         gp.schedlink,
			Waitsince:         gp.waitsince,
			Waitreason:        gp.waitreason,
			Preempt:           gp.preempt,
			PreemptStop:       gp.preemptStop,
			PreemptShrink:     gp.preemptShrink,
			AsyncSafePoint:    gp.asyncSafePoint,
			Paniconfault:      gp.paniconfault,
			Gcscandone:        gp.gcscandone,
			Throwsplit:        gp.throwsplit,
			ActiveStackChans:  gp.activeStackChans,
			ParkingOnChan:     gp.parkingOnChan,
			InMarkAssist:      gp.inMarkAssist,
			Coroexit:          gp.coroexit,
			Raceignore:        gp.raceignore,
			Nocgocallback:     gp.nocgocallback,
			Tracking:          gp.tracking,
			TrackingSeq:       gp.trackingSeq,
			TrackingStamp:     gp.trackingStamp,
			RunnableTime:      gp.runnableTime,
			Lockedm:           gp.lockedm,
			FipsIndicator:     gp.fipsIndicator,
			Sig:               gp.sig,
			Writebuf:          gp.writebuf,
			Sigcode0:          gp.sigcode0,
			Sigcode1:          gp.sigcode1,
			Sigpc:             gp.sigpc,
			ParentGoid:        gp.parentGoid,
			Gopc:              gp.gopc,
			Ancestors:         gp.ancestors,
			Startpc:           gp.startpc,
			Racectx:           gp.racectx,
			Waiting:           gp.waiting,
			CgoCtxt:           gp.cgoCtxt,
			Labels:            gp.labels,
			Timer:             gp.timer,
			SleepWhen:         gp.sleepWhen,
			SelectDone:        gp.selectDone,
			GoroutineProfiled: gp.goroutineProfiled,
			Coroarg:           gp.coroarg,
			SyncGroup:         gp.syncGroup,
			Trace:             gp.trace,
			GcAssistBytes:     gp.gcAssistBytes,
			Annotations: gp.annotations,
		}
		all = append(all, g)
	})
	return all
}

func ForEachM() []*M {
	var all []*M
	lock(&sched.lock)
	defer unlock(&sched.lock)
	for mp := allm; mp != nil; mp = mp.alllink {
		m := &M{
			G0:              mp.g0,
			Morebuf:         mp.morebuf,
			Divmod:          mp.divmod,
			Procid:          mp.procid,
			Gsignal:         mp.gsignal,
			GoSigStack:      mp.goSigStack,
			Sigmask:         mp.sigmask,
			TLS:             mp.tls,
			Mstartfn:        mp.mstartfn,
			Curg:            mp.curg,
			Caughtsig:       mp.caughtsig,
			P:               mp.p,
			Nextp:           mp.nextp,
			Oldp:            mp.oldp,
			ID:              mp.id,
			Mallocing:       mp.mallocing,
			Throwing:        mp.throwing,
			Preemptoff:      mp.preemptoff,
			Locks:           mp.locks,
			Dying:           mp.dying,
			Profilehz:       mp.profilehz,
			Spinning:        mp.spinning,
			Blocked:         mp.blocked,
			NewSigstack:     mp.newSigstack,
			Printlock:       mp.printlock,
			Incgo:           mp.incgo,
			Isextra:         mp.isextra,
			IsExtraInC:      mp.isExtraInC,
			IsExtraInSig:    mp.isExtraInSig,
			FreeWait:        mp.freeWait,
			Needextram:      mp.needextram,
			G0StackAccurate: mp.g0StackAccurate,
			Traceback:       mp.traceback,
			Ncgocall:        mp.ncgocall,
			Ncgo:            mp.ncgo,
			CgoCallersUse:   mp.cgoCallersUse,
			CgoCallers:      mp.cgoCallers,
			Park:            mp.park,
			Alllink:         mp.alllink,
			Schedlink:       mp.schedlink,
			Lockedg:         mp.lockedg,
			Createstack:     mp.createstack,
			LockedExt:       mp.lockedExt,
			LockedInt:       mp.lockedInt,
			MWaitList:       mp.mWaitList,

			MLockProfile: mp.mLockProfile,
			ProfStack:    mp.profStack,

			Waitunlockf:          mp.waitunlockf,
			Waitlock:             mp.waitlock,
			WaitTraceSkip:        mp.waitTraceSkip,
			WaitTraceBlockReason: mp.waitTraceBlockReason,

			Syscalltick: mp.syscalltick,
			Freelink:    mp.freelink,
			Trace:       mp.trace,

			Libcall:    mp.libcall,
			Libcallpc:  mp.libcallpc,
			Libcallsp:  mp.libcallsp,
			Libcallg:   mp.libcallg,
			Winsyscall: mp.winsyscall,

			VdsoSP: mp.vdsoSP,
			VdsoPC: mp.vdsoPC,

			PreemptGen: mp.preemptGen,

			SignalPending: mp.signalPending,

			PcvalueCache: mp.pcvalueCache,

			DlogPerM: mp.dlogPerM,

			MOS: mp.mOS,

			Chacha8:   mp.chacha8,
			Cheaprand: mp.cheaprand,

			LocksHeldLen: mp.locksHeldLen,
			LocksHeld:    mp.locksHeld,
		}
		all = append(all, m)
	}
	return all
}

func ForEachP() []*P {
	all := make([]*P, 0, gomaxprocs)
	lock(&sched.lock)
	defer unlock(&sched.lock)
	for _, p := range allp {
		pp := &P{
			ID:                    p.id,
			Status:                p.status,
			Link:                  p.link,
			Schedtick:             p.schedtick,
			Syscalltick:           p.syscalltick,
			Sysmontick:            p.sysmontick,
			M:                     p.m,
			Mcache:                p.mcache,
			Pcache:                p.pcache,
			Raceprocctx:           p.raceprocctx,
			Deferpool:             p.deferpool,
			Deferpoolbuf:          p.deferpoolbuf,
			Goidcache:             p.goidcache,
			Goidcacheend:          p.goidcacheend,
			Runqhead:              p.runqhead,
			Runqtail:              p.runqtail,
			Runq:                  p.runq,
			Runnext:               p.runnext,
			GFree:                 p.gFree,
			Sudogcache:            p.sudogcache,
			Sudogbuf:              p.sudogbuf,
			Mspancache:            p.mspancache,
			PinnerCache:           p.pinnerCache,
			Trace:                 p.trace,
			Palloc:                p.palloc,
			GcAssistTime:          p.gcAssistTime,
			GcFractionalMarkTime:  p.gcFractionalMarkTime,
			LimiterEvent:          p.limiterEvent,
			GcMarkWorkerMode:      p.gcMarkWorkerMode,
			GcMarkWorkerStartTime: p.gcMarkWorkerStartTime,
			Gcw:                   p.gcw,
			WbBuf:                 p.wbBuf,
			RunSafePointFn:        p.runSafePointFn,
			StatsSeq:              p.statsSeq,
			Timers:                p.timers,
			MaxStackScanDelta:     p.maxStackScanDelta,
			ScannedStackSize:      p.scannedStackSize,
			ScannedStacks:         p.scannedStacks,
			Preempt:               p.preempt,
			GcStopTime:            p.gcStopTime,
		}
		all = append(all, pp)
	}
	return all
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
