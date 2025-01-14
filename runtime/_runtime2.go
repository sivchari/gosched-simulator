package runtime

var WaitReasonStrings = [...]string{
	waitReasonZero:                  "",
	waitReasonGCAssistMarking:       "GC assist marking",
	waitReasonIOWait:                "IO wait",
	waitReasonChanReceiveNilChan:    "chan receive (nil chan)",
	waitReasonChanSendNilChan:       "chan send (nil chan)",
	waitReasonDumpingHeap:           "dumping heap",
	waitReasonGarbageCollection:     "garbage collection",
	waitReasonGarbageCollectionScan: "garbage collection scan",
	waitReasonPanicWait:             "panicwait",
	waitReasonSelect:                "select",
	waitReasonSelectNoCases:         "select (no cases)",
	waitReasonGCAssistWait:          "GC assist wait",
	waitReasonGCSweepWait:           "GC sweep wait",
	waitReasonGCScavengeWait:        "GC scavenge wait",
	waitReasonChanReceive:           "chan receive",
	waitReasonChanSend:              "chan send",
	waitReasonFinalizerWait:         "finalizer wait",
	waitReasonForceGCIdle:           "force gc (idle)",
	waitReasonSemacquire:            "semacquire",
	waitReasonSleep:                 "sleep",
	waitReasonSyncCondWait:          "sync.Cond.Wait",
	waitReasonSyncMutexLock:         "sync.Mutex.Lock",
	waitReasonSyncRWMutexRLock:      "sync.RWMutex.RLock",
	waitReasonSyncRWMutexLock:       "sync.RWMutex.Lock",
	waitReasonSyncWaitGroupWait:     "sync.WaitGroup.Wait",
	waitReasonTraceReaderBlocked:    "trace reader (blocked)",
	waitReasonWaitForGCCycle:        "wait for GC cycle",
	waitReasonGCWorkerIdle:          "GC worker (idle)",
	waitReasonGCWorkerActive:        "GC worker (active)",
	waitReasonPreempted:             "preempted",
	waitReasonDebugCall:             "debug call",
	waitReasonGCMarkTermination:     "GC mark termination",
	waitReasonStoppingTheWorld:      "stopping the world",
	waitReasonFlushProcCaches:       "flushing proc caches",
	waitReasonTraceGoroutineStatus:  "trace goroutine status",
	waitReasonTraceProcStatus:       "trace proc status",
	waitReasonPageTraceFlush:        "page trace flush",
	waitReasonCoroutine:             "coroutine",
	waitReasonGCWeakToStrongWait:    "GC weak to strong wait",
	waitReasonSynctestRun:           "synctest.Run",
	waitReasonSynctestWait:          "synctest.Wait",
	waitReasonSynctestChanReceive:   "chan receive (synctest)",
	waitReasonSynctestChanSend:      "chan send (synctest)",
	waitReasonSynctestSelect:        "select (synctest)",
}

type GFree struct {
	GList gList
	N     int32
}

type g struct {
	stack       stack
	stackguard0 uintptr
	stackguard1 uintptr

	_panic       *_panic
	_defer       *_defer
	m            *m
	sched        gobuf
	syscallsp    uintptr
	syscallpc    uintptr
	syscallbp    uintptr
	stktopsp     uintptr
	param        unsafe.Pointer
	atomicstatus atomic.Uint32
	stackLock    uint32
	goid         uint64
	schedlink    guintptr
	waitsince    int64
	waitreason   waitReason

	preempt       bool
	preemptStop   bool
	preemptShrink bool

	asyncSafePoint bool

	paniconfault     bool
	gcscandone       bool
	throwsplit       bool
	activeStackChans bool
	parkingOnChan    atomic.Bool
	inMarkAssist     bool
	coroexit         bool

	raceignore    int8
	nocgocallback bool
	tracking      bool
	trackingSeq   uint8
	trackingStamp int64
	runnableTime  int64
	lockedm       muintptr
	fipsIndicator uint8
	sig           uint32
	writebuf      []byte
	sigcode0      uintptr
	sigcode1      uintptr
	sigpc         uintptr
	parentGoid    uint64
	gopc          uintptr
	ancestors     *[]ancestorInfo
	startpc       uintptr
	racectx       uintptr
	waiting       *sudog
	cgoCtxt       []uintptr
	labels        unsafe.Pointer
	timer         *timer
	sleepWhen     int64
	selectDone    atomic.Uint32

	goroutineProfiled goroutineProfileStateHolder

	coroarg   *coro
	syncGroup *synctestGroup

	trace gTraceState

	gcAssistBytes int64

	// annotations are that gosched-simulator attached.
	annotations []string
}

type G struct {
	Stack       stack
	Stackguard0 uintptr
	Stackguard1 uintptr

	Panic        *_panic
	Defer        *_defer
	M            *m
	Sched        gobuf
	Syscallsp    uintptr
	Syscallpc    uintptr
	Syscallbp    uintptr
	Stktopsp     uintptr
	Param        unsafe.Pointer
	Atomicstatus atomic.Uint32
	StackLock    uint32
	Goid         uint64
	Schedlink    guintptr
	Waitsince    int64
	Waitreason   waitReason

	Preempt       bool
	PreemptStop   bool
	PreemptShrink bool

	AsyncSafePoint bool

	Paniconfault     bool
	Gcscandone       bool
	Throwsplit       bool
	ActiveStackChans bool
	ParkingOnChan    atomic.Bool
	InMarkAssist     bool
	Coroexit         bool

	Raceignore    int8
	Nocgocallback bool
	Tracking      bool
	TrackingSeq   uint8
	TrackingStamp int64
	RunnableTime  int64
	Lockedm       muintptr
	FipsIndicator uint8
	Sig           uint32
	Writebuf      []byte
	Sigcode0      uintptr
	Sigcode1      uintptr
	Sigpc         uintptr
	ParentGoid    uint64
	Gopc          uintptr
	Ancestors     *[]ancestorInfo
	Startpc       uintptr
	Racectx       uintptr
	Waiting       *sudog
	CgoCtxt       []uintptr
	Labels        unsafe.Pointer
	Timer         *timer
	SleepWhen     int64
	SelectDone    atomic.Uint32

	GoroutineProfiled goroutineProfileStateHolder

	Coroarg   *coro
	SyncGroup *synctestGroup

	Trace gTraceState

	GcAssistBytes int64

	// Annotations are that gosched-simulator attached.
	Annotations []string
}

type M struct {
	G0      *g
	Morebuf gobuf
	Divmod  uint32
	_       uint32

	Procid          uint64
	Gsignal         *g
	GoSigStack      gsignalStack
	Sigmask         sigset
	TLS             [tlsSlots]uintptr
	Mstartfn        func()
	Curg            *g
	Caughtsig       guintptr
	P               puintptr
	Nextp           puintptr
	Oldp            puintptr
	ID              int64
	Mallocing       int32
	Throwing        throwType
	Preemptoff      string
	Locks           int32
	Dying           int32
	Profilehz       int32
	Spinning        bool
	Blocked         bool
	NewSigstack     bool
	Printlock       int8
	Incgo           bool
	Isextra         bool
	IsExtraInC      bool
	IsExtraInSig    bool
	FreeWait        atomic.Uint32
	Needextram      bool
	G0StackAccurate bool
	Traceback       uint8
	Ncgocall        uint64
	Ncgo            int32
	CgoCallersUse   atomic.Uint32
	CgoCallers      *cgoCallers
	Park            note
	Alllink         *m
	Schedlink       muintptr
	Lockedg         guintptr
	Createstack     [32]uintptr
	LockedExt       uint32
	LockedInt       uint32
	MWaitList       mWaitList

	MLockProfile mLockProfile
	ProfStack    []uintptr

	Waitunlockf          func(*g, unsafe.Pointer) bool
	Waitlock             unsafe.Pointer
	WaitTraceSkip        int
	WaitTraceBlockReason traceBlockReason

	Syscalltick uint32
	Freelink    *m
	Trace       mTraceState

	Libcall    libcall
	Libcallpc  uintptr
	Libcallsp  uintptr
	Libcallg   guintptr
	Winsyscall winlibcall

	VdsoSP uintptr
	VdsoPC uintptr

	PreemptGen atomic.Uint32

	SignalPending atomic.Uint32

	PcvalueCache pcvalueCache

	DlogPerM dlogPerM

	MOS mOS

	Chacha8   chacha8rand.State
	Cheaprand uint64

	LocksHeldLen int
	LocksHeld    [10]heldLockInfo

	_ [goexperiment.SpinbitMutexInt * 700 * (2 - goarch.PtrSize/4)]byte
}

type P struct {
	ID          int32
	Status      uint32
	Link        puintptr
	Schedtick   uint32
	Syscalltick uint32
	Sysmontick  sysmontick
	M           muintptr
	Mcache      *mcache
	Pcache      pageCache
	Raceprocctx uintptr

	Deferpool    []*_defer
	Deferpoolbuf [32]*_defer

	Goidcache    uint64
	Goidcacheend uint64

	Runqhead uint32
	Runqtail uint32
	Runq     [256]guintptr
	Runnext  guintptr

	GFree GFree

	Sudogcache []*sudog
	Sudogbuf   [128]*sudog

	Mspancache struct {
		len int
		buf [128]*mspan
	}

	PinnerCache *pinner

	Trace pTraceState

	Palloc persistentAlloc

	GcAssistTime         int64
	GcFractionalMarkTime int64

	LimiterEvent limiterEvent

	GcMarkWorkerMode      gcMarkWorkerMode
	GcMarkWorkerStartTime int64

	Gcw gcWork

	WbBuf wbBuf

	RunSafePointFn uint32

	StatsSeq atomic.Uint32

	Timers timers

	MaxStackScanDelta int64

	ScannedStackSize uint64
	ScannedStacks    uint64

	Preempt bool

	GcStopTime int64
}

func (gptr guintptr) Ptr() *G {
	gp := gptr.ptr()
	return castg(gp)
}

func (mptr muintptr) Ptr() *M {
	mp := mptr.ptr()
	return castm(mp)
}

func (pptr puintptr) Ptr() *P {
	pp := pptr.ptr()
	return castp(pp)
}

func (g *g) GoID() uint64 {
	if g == nil {
		return 0
	}
	return g.goid
}

func (m *m) ID() int64 {
	if m == nil {
		return 0
	}
	return m.id
}

func (m *m) ProcID() uint64 {
	if m == nil {
		return 0
	}
	return m.procid
}

func (p *p) ID() int32 {
	if p == nil {
		return 0
	}
	return p.id
}

func (gList gList) Pop() *G {
	gp := gList.pop()
	return castg(gp)
}

func (gList *gList) Empty() bool {
	return gList.empty()
}

func (gQueue *gQueue) Pop() *G {
	gp := gQueue.pop()
	return castg(gp)
}

func (gQueue *gQueue) Empty() bool {
	return gQueue.empty()
}

func castg(gp *g) *G {
	if gp == nil {
		return nil
	}
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
		Annotations:       gp.annotations,
	}
	return g
}

func castm(mp *m) *M {
	if mp == nil {
		return nil
	}
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
	return m
}

func castp(pp *p) *P {
	if pp == nil {
		return nil
	}
	p := &P{
		ID:           pp.id,
		Status:       pp.status,
		Link:         pp.link,
		Schedtick:    pp.schedtick,
		Syscalltick:  pp.syscalltick,
		Sysmontick:   pp.sysmontick,
		M:            pp.m,
		Mcache:       pp.mcache,
		Pcache:       pp.pcache,
		Raceprocctx:  pp.raceprocctx,
		Deferpool:    pp.deferpool,
		Deferpoolbuf: pp.deferpoolbuf,
		Goidcache:    pp.goidcache,
		Goidcacheend: pp.goidcacheend,
		Runqhead:     pp.runqhead,
		Runqtail:     pp.runqtail,
		Runq:         pp.runq,
		Runnext:      pp.runnext,
		GFree: GFree{
			GList: pp.gFree.gList,
			N:     pp.gFree.n,
		},
		Sudogcache:            pp.sudogcache,
		Sudogbuf:              pp.sudogbuf,
		Mspancache:            pp.mspancache,
		PinnerCache:           pp.pinnerCache,
		Trace:                 pp.trace,
		Palloc:                pp.palloc,
		GcAssistTime:          pp.gcAssistTime,
		GcFractionalMarkTime:  pp.gcFractionalMarkTime,
		LimiterEvent:          pp.limiterEvent,
		GcMarkWorkerMode:      pp.gcMarkWorkerMode,
		GcMarkWorkerStartTime: pp.gcMarkWorkerStartTime,
		Gcw:                   pp.gcw,
		WbBuf:                 pp.wbBuf,
		RunSafePointFn:        pp.runSafePointFn,
		StatsSeq:              pp.statsSeq,
		Timers:                pp.timers,
		MaxStackScanDelta:     pp.maxStackScanDelta,
		ScannedStackSize:      pp.scannedStackSize,
		ScannedStacks:         pp.scannedStacks,
		Preempt:               pp.preempt,
		GcStopTime:            pp.gcStopTime,
	}
	return p
}

type Schedt struct {
	Goidgen   atomic.Uint64
	Lastpoll  atomic.Int64
	PollUntil atomic.Int64

	Lock mutex

	Midle        muintptr
	Nmidle       int32
	Nmidlelocked int32
	Mnext        int64
	Maxmcount    int32
	Nmsys        int32
	Nmfreed      int64

	Ngsys atomic.Int32

	Pidle        puintptr
	Npidle       atomic.Int32
	Nmspinning   atomic.Int32
	Needspinning atomic.Uint32

	Runq     gQueue
	Runqsize int32

	Disable struct {
		user     bool
		runnable gQueue
		n        int32
	}

	GFree GlobGFree

	Sudoglock  mutex
	Sudogcache *sudog

	Deferlock mutex
	Deferpool *_defer

	Freem *m

	Gcwaiting  atomic.Bool
	Stopwait   int32
	Stopnote   note
	Sysmonwait atomic.Bool
	Sysmonnote note

	SafePointFn   func(*p)
	SafePointWait int32
	SafePointNote note

	Profilehz int32

	Procresizetime int64
	Totaltime      int64

	Sysmonlock mutex

	TimeToRun timeHistogram

	IdleTime atomic.Int64

	TotalMutexWaitTime atomic.Int64

	StwStoppingTimeGC    timeHistogram
	StwStoppingTimeOther timeHistogram

	StwTotalTimeGC    timeHistogram
	StwTotalTimeOther timeHistogram

	TotalRuntimeLockWaitTime atomic.Int64
}

func castsched() *Schedt {
	return &Schedt{
		Goidgen:   sched.goidgen,
		Lastpoll:  sched.lastpoll,
		PollUntil: sched.pollUntil,

		Lock: sched.lock,

		Midle:        sched.midle,
		Nmidle:       sched.nmidle,
		Nmidlelocked: sched.nmidlelocked,
		Mnext:        sched.mnext,
		Maxmcount:    sched.maxmcount,
		Nmsys:        sched.nmsys,
		Nmfreed:      sched.nmfreed,

		Ngsys: sched.ngsys,

		Pidle:        sched.pidle,
		Npidle:       sched.npidle,
		Nmspinning:   sched.nmspinning,
		Needspinning: sched.needspinning,

		Runq:     sched.runq,
		Runqsize: sched.runqsize,

		Disable: sched.disable,

		GFree: GlobGFree{
			Lock:    sched.gFree.lock,
			Stack:   sched.gFree.stack,
			NoStack: sched.gFree.noStack,
			N:       sched.gFree.n,
		},

		Sudoglock:  sched.sudoglock,
		Sudogcache: sched.sudogcache,

		Deferlock: sched.deferlock,
		Deferpool: sched.deferpool,

		Freem: sched.freem,

		Gcwaiting:  sched.gcwaiting,
		Stopwait:   sched.stopwait,
		Stopnote:   sched.stopnote,
		Sysmonwait: sched.sysmonwait,
		Sysmonnote: sched.sysmonnote,

		SafePointFn:   sched.safePointFn,
		SafePointWait: sched.safePointWait,
		SafePointNote: sched.safePointNote,

		Profilehz: sched.profilehz,

		Procresizetime: sched.procresizetime,
		Totaltime:      sched.totaltime,

		Sysmonlock: sched.sysmonlock,

		TimeToRun: sched.timeToRun,

		IdleTime: sched.idleTime,

		TotalMutexWaitTime: sched.totalMutexWaitTime,

		StwStoppingTimeGC:    sched.stwStoppingTimeGC,
		StwStoppingTimeOther: sched.stwStoppingTimeOther,

		StwTotalTimeGC:    sched.stwTotalTimeGC,
		StwTotalTimeOther: sched.stwTotalTimeOther,

		TotalRuntimeLockWaitTime: sched.totalRuntimeLockWaitTime,
	}
}

type GlobGFree struct {
	Lock    mutex
	Stack   gList
	NoStack gList
	N       int32
}
