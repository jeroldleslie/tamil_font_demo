
export interface Marker {
	lat: number;
	lng: number;
	label?: string;
	title?: string;
  id: string;
}

export interface School {
	name: string;
	city: string;
	lat: number;
	lng: number;
	id: number;
	percs: percreq[];
	gpas:	gpareq[];
	images: string[];
	acceptance_rate: string,
	total_students: number,
	male_ratio: string,
	female_ratio: string,
	student_population: number,
	freshman_class_size: number,
	avg_financial_aid_package: string
}

export interface percreq {
	perc: number;
	act: number;
	sat: number;
}

export interface gpareq {
	gpa: number;
	act: number;
	sat: number;
}


