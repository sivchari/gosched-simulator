export interface G {
	goid: number;
	waitreason: string;
	annotations: string[];
	status: string;
}

export interface M {
	procid: number;
	curg: G;
	id: number;
}

export interface P {
	id: number;
	m: M;
	runq: G[];
	gfree: G[];
}

export interface Scheduler {
	runq: G[];
	stack: G[];
	noStack: G[];
}
