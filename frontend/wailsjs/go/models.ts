export namespace models {
	
	export class Profile {
	    twitchKey: string;
	    psnNpsso: string;
	
	    static createFrom(source: any = {}) {
	        return new Profile(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitchKey = source["twitchKey"];
	        this.psnNpsso = source["psnNpsso"];
	    }
	}
	export class ReadProfileOutput {
	    userProfile: Profile;
	
	    static createFrom(source: any = {}) {
	        return new ReadProfileOutput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.userProfile = this.convertValues(source["userProfile"], Profile);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
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
	export class UpdateProfileInput {
	    twitchKey: string;
	    psnNpsso: string;
	
	    static createFrom(source: any = {}) {
	        return new UpdateProfileInput(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.twitchKey = source["twitchKey"];
	        this.psnNpsso = source["psnNpsso"];
	    }
	}

}

