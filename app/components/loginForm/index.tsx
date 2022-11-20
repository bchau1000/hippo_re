import React, { Dispatch, SetStateAction } from "react";
import DefaultInput from "../defaultInput";

type LoginFormProps = {
    className: string, 
    setEmail: Dispatch<SetStateAction<string | undefined>>,
    setPassword: Dispatch<SetStateAction<string | undefined>>,
    onSubmit: () => any
} & typeof defaultProps;

const defaultProps = {
    className: "login-form",
    onSubmit: () => console.warn("No submit function was provided for LoginForm component")
}

const LoginForm = ({
    className,
    setEmail,
    setPassword,
    onSubmit,
}:LoginFormProps) => {
    const emailInputClass:string = `${className}-email-input`;
    const passwordInputClass:string = `${className}-password-input`;
    const submitButtonClass:string = `${className}-submit`;

    return (
        <>
            <form 
                className={className} 
                onSubmit={(event) => {
                    event.preventDefault();
                    onSubmit();
                }}
            >
                <DefaultInput
                    className={emailInputClass}
                    type={"email"}
                    label={"Email"}
                    onChange={(event) => setEmail(event.currentTarget.value)}
                />
                <DefaultInput
                    className={passwordInputClass}
                    type={"password"}
                    label={"Password"}
                    onChange={(event) => setPassword(event.currentTarget.value)}
                />
                <button 
                    className={submitButtonClass}
                    type="submit"
                >
                    Login
                </button>
            </form>
        </>
    );
}

LoginForm.defaultProps = defaultProps;
export default LoginForm;