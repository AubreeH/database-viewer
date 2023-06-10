export enum EShadowWeight {
	Light = "var(--box-shadow-light)",
	Medium = "var(--box-shadow-med)",
	Heavy = "var(--box-shadow-heavy)",
}

export interface IShadowProperties {
	weight: EShadowWeight | string;
	size: number;
}
