import { ButtonHTMLAttributes, InputHTMLAttributes, TextareaHTMLAttributes } from "react";

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