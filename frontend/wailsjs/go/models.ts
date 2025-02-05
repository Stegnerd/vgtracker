export namespace config {
	
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
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitch = this.convertValues(source["twitch"], Twitch);
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
	    twitch: config.Twitch;
	
	    static createFrom(source: any = {}) {
	        return new UpdateConfigInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitch = this.convertValues(source["twitch"], config.Twitch);
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

