// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function AddBookToCollection(arg1:number,arg2:number):Promise<void>;

export function CreateCollection(arg1:models.CollectionInfo):Promise<void>;

export function DeleteBookToCollection(arg1:number,arg2:number):Promise<void>;

export function DeleteCollection(arg1:number):Promise<void>;

export function GetCollection(arg1:string):Promise<models.Collection>;

export function GetCollectionById(arg1:number):Promise<models.Collection>;

export function UpdateCollection(arg1:number,arg2:models.CollectionInfo):Promise<void>;
