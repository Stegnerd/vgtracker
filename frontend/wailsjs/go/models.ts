export namespace models {
	
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

