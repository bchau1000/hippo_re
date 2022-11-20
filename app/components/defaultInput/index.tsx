import { FormEvent } from "react";

type DefaultInputProps = {
    className:string,
    label:string,
    type:string,
    onChange:(event:FormEvent<HTMLInputElement>) => any
} & typeof defaultProps;

const defaultProps = {
    baseClass: "base-default-input",
    className: "default-input",
    label: "",
    type: "text",
};

const DefaultInput = ({
    baseClass,
    className, 
    label, 
    type,
    onChange,
}:DefaultInputProps) => {
    const inputClassName = className;
    const labelClassName = `${className}-label`;

    return (
        <>
            {label != null && label.length > 0 
                ? (
                    <label 
                        className={`${baseClass} ${labelClassName}`} 
                        htmlFor={className}
                        
                    >
                        {label}
                    </label>
                ):(
                    <>
                    </>
                )
            }
            <input 
                className={`${baseClass} ${inputClassName}`} 
                type={type}
                onChange={(event) => onChange(event)}
            />
        </>
    );
}

DefaultInput.defaultProps = defaultProps;
export default DefaultInput;