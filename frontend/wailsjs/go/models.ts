export namespace models {
	
	export interface BookAuthor {
	    name: string;
	    birth_year: number;
	    death_year: number;
	}
	export interface Book {
	    id: number;
	    title: string;
	    authors: BookAuthor[];
	    subjects: string[];
	    bookshelves: string[];
	    languages: string[];
	    copyright: boolean;
	    media_type: string;
	    formats: {[key: string]: string};
	    download_count: number;
	}
	
	
	export interface Collection {
	    id: number;
	    name: string;
	    description: string;
	    documents: number[];
	    owner_id: number;
	    category: string;
	    public: boolean;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	}
	export interface CollectionInfo {
	    name: string;
	    description: string;
	    documents: number[];
	    owner_id: number;
	    category: string;
	    public: boolean;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	}
	export interface Err {
	    error: any;
	    message: string;
	    status: number;
	}
	export interface Gutendex {
	    count: number;
	    next: string;
	    previous: string;
	    results: Book[];
	}

}

