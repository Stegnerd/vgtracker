export namespace config {
	
	export class ThemeSettings {
	    preset: models.PresetConfig;
	    primaryColor: models.PaletteColor;
	    surfaceColor: models.SufaceColor;
	    IsDarkTheme: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ThemeSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.preset = source["preset"];
	        this.primaryColor = source["primaryColor"];
	        this.surfaceColor = source["surfaceColor"];
	        this.IsDarkTheme = source["IsDarkTheme"];
	    }
	}
	export class Steam {
	    steamID: string;
	    apiKey: string;
	
	    static createFrom(source: any = {}) {
	        return new Steam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.steamID = source["steamID"];
	        this.apiKey = source["apiKey"];
	    }
	}
	export class Twitch {
	    clientID: string;
	    clientSecret: string;
	
	    static createFrom(source: any = {}) {
	        return new Twitch(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.clientID = source["clientID"];
	        this.clientSecret = source["clientSecret"];
	    }
	}
	export class Config {
	    twitch: Twitch;
	    steam: Steam;
	    theme: ThemeSettings;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitch = this.convertValues(source["twitch"], Twitch);
	        this.steam = this.convertValues(source["steam"], Steam);
	        this.theme = this.convertValues(source["theme"], ThemeSettings);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

export namespace controllers {
	
	export class UpdateConfigInput {
	    twitch?: config.Twitch;
	    steam?: config.Steam;
	    preset?: models.PresetConfig;
	    primaryColor?: models.PaletteColor;
	    surfaceColor?: models.SufaceColor;
	    IsDarkTheme?: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UpdateConfigInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitch = this.convertValues(source["twitch"], config.Twitch);
	        this.steam = this.convertValues(source["steam"], config.Steam);
	        this.preset = source["preset"];
	        this.primaryColor = source["primaryColor"];
	        this.surfaceColor = source["surfaceColor"];
	        this.IsDarkTheme = source["IsDarkTheme"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace gamedetails {
	
	export class GetGameDetailOutput {
	    id: number;
	    igdbID: number;
	    isOwned: boolean;
	    isBeaten: boolean;
	    isWishlisted: boolean;
	    name: string;
	    thumbnailURL: string;
	    coverURL: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new GetGameDetailOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.igdbID = source["igdbID"];
	        this.isOwned = source["isOwned"];
	        this.isBeaten = source["isBeaten"];
	        this.isWishlisted = source["isWishlisted"];
	        this.name = source["name"];
	        this.thumbnailURL = source["thumbnailURL"];
	        this.coverURL = source["coverURL"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class UpsertGameDetailInput {
	    id?: number;
	    igdbID?: number;
	    isOwned?: boolean;
	    isBeaten?: boolean;
	    isWishlisted?: boolean;
	    name?: string;
	    thumbnailURL?: string;
	    coverURL?: string;
	
	    static createFrom(source: any = {}) {
	        return new UpsertGameDetailInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.igdbID = source["igdbID"];
	        this.isOwned = source["isOwned"];
	        this.isBeaten = source["isBeaten"];
	        this.isWishlisted = source["isWishlisted"];
	        this.name = source["name"];
	        this.thumbnailURL = source["thumbnailURL"];
	        this.coverURL = source["coverURL"];
	    }
	}

}

export namespace igdb {
	
	export class VGTGame {
	    id: number;
	    name: string;
	    gameType: string;
	    genres: string[];
	    platforms: string[];
	    thumbnailURL: string;
	    coverURL: string;
	    firstReleaseYear: number;
	    summary: string;
	
	    static createFrom(source: any = {}) {
	        return new VGTGame(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.gameType = source["gameType"];
	        this.genres = source["genres"];
	        this.platforms = source["platforms"];
	        this.thumbnailURL = source["thumbnailURL"];
	        this.coverURL = source["coverURL"];
	        this.firstReleaseYear = source["firstReleaseYear"];
	        this.summary = source["summary"];
	    }
	}
	export class VGTSearchResults {
	    items: VGTGame[];
	
	    static createFrom(source: any = {}) {
	        return new VGTSearchResults(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], VGTGame);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace models {
	
	export enum PaletteColor {
	    noir = "noir",
	    emerald = "emerald",
	    green = "green",
	    lime = "lime",
	    orange = "orange",
	    amber = "amber",
	    yellow = "yellow",
	    teal = "teal",
	    cyan = "cyan",
	    sky = "sky",
	    blue = "blue",
	    indigo = "indigo",
	    violet = "violet",
	    purple = "purple",
	    fuschia = "fuchsia",
	    pink = "pink",
	    rose = "rose",
	}
	export enum SufaceColor {
	    slate = "slate",
	    gray = "gray",
	    zinc = "zinc",
	    neutral = "neutral",
	    stone = "stone",
	    soho = "soho",
	    viva = "viva",
	    ocean = "ocean",
	}
	export enum PresetConfig {
	    Aura = "Aura",
	    Lara = "Lara",
	    Nora = "Nora",
	}

}

export namespace twitch {
	
	export class AccessTokenResponse {
	    access_token: string;
	    expires_in: number;
	    token_type: string;
	
	    static createFrom(source: any = {}) {
	        return new AccessTokenResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.access_token = source["access_token"];
	        this.expires_in = source["expires_in"];
	        this.token_type = source["token_type"];
	    }
	}

}

