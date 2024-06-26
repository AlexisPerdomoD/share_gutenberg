import { ZodType, z } from 'zod';
import { models as m} from '../../wailsjs/go/models';

// Esquema de Zod para el tipo Book

const bookAuthorSchema: ZodType<m.BookAuthor> = z.object({
    name: z.string().default(''),
    birth_year: z.number().default(0),
    death_year: z.number().default(0),
});
export const bookSchema: ZodType<m.Book> = z.object({
    id: z.number().min(1).default(0),
    title: z.string().default(''),
    authors: z.array(bookAuthorSchema).default([]),
    subjects: z.array(z.string()).default([]),
    bookshelves: z.array(z.string()).default([]),
    languages: z.array(z.string()).default([]),
    copyright: z.boolean().default(false),
    media_type: z.string().default(''),
    formats: z.record(z.string()).default({}),
    download_count: z.number().default(0),
})
/*
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
	}*/
export const PASS_REGEX = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$/

export const userSchema:ZodType<m.UserInfo> = z.object({
    name: z.string().max(50).default(""),
    password:z.string().min(8).regex(PASS_REGEX),
    email:z.string().email(),
    role:z.enum(["user", "admin"]).default("user"),
    created_at: z.date().default(new Date()),
    updated_at:z.date().default(new Date())

})

// export interface UserInfo {
// 	    name: string;
// 	    email: string;
// 	    password: string;
// 	    role: string;
// 	    // Go type: time
// 	    created_at: any;
// 	    // Go type: time
// 	    updated_at: any;
// 	}
const validate = <T>(schema: ZodType<T>, data: unknown): T | undefined => {
    try {
        return schema.parse(data) as T
    } catch (error) {
        console.error('Error al validar los datos:', error);
        return undefined;
    }
}
export default validate
