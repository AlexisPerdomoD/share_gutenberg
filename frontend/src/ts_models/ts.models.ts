import { ButtonHTMLAttributes, InputHTMLAttributes, LabelHTMLAttributes, TextareaHTMLAttributes } from "react";
export interface Current{
    username:string,
    name: string, 
    role: "user" | "admin"
}
export interface CurrentContext{
    current:Current | null,
    setCurrent:Function
}
export interface Queries{
    [key: string]: string 
    category:string,
    page:string,
    topic:string,
    ids:string,
    copyright:string,
    languages:string,
    search:string
}

/* */
export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
    text:string;
}
export interface InputProps extends InputHTMLAttributes<HTMLInputElement>{}
export interface TextAreaProps extends TextareaHTMLAttributes<HTMLTextAreaElement> {}


export interface LabelProps extends LabelHTMLAttributes<HTMLLabelElement>{}
