// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {models} from '../models';

export function AddCollection(arg1:number,arg2:number):Promise<void>;

export function CreateUser(arg1:models.UserInfo):Promise<void>;

export function DeleteUser(arg1:number):Promise<void>;

export function GetUser(arg1:number):Promise<models.User|models.Err>;

export function GetUserByEmail(arg1:string):Promise<models.User|models.Err>;

export function GetUserCollections(arg1:number):Promise<models.UserCollections>;

export function RemoveCollection(arg1:number,arg2:number):Promise<void>;

export function UpdateUser(arg1:number,arg2:models.UserInfo):Promise<void>;