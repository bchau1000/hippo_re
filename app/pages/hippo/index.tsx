import styles from '../../styles/Home.module.css'
import React, { useState, useEffect} from 'react'
import { PingModel } from '../../models'
import { pingServer, loginUser, authUser } from '../api';
import { AuthError, User, UserCredential } from "firebase/auth";
import { LoginForm } from '../../components';

const login = async (emailOpt:string|undefined, passwordOpt:string|undefined):Promise<boolean> => {
    let email:string = emailOpt || '';
    let password:string = passwordOpt || '';

    return loginUser(email, password)
        .then(async (response:Response):Promise<boolean> => {
            if (response.status === 200) {
                console.log(await response.json());
            } else {
                throw new Error("Invalid username or password")
            }
            return true;
        })
        .catch((error:AuthError) => {
            console.log(error)
            return false;
        })
        .catch((error:any) => {
            console.log(error)
            return false;
        });
}

const auth = async ():Promise<boolean> => {
    const isLoggedIn:boolean = await authUser();
    return isLoggedIn;
}

export default function Home() {
    const [ping, setPing] = useState<PingModel | undefined>();
    const [loggedIn, setLoggedIn] = useState<boolean | undefined>(true);
    const [loading, setLoading] = useState<boolean | undefined>(true);
    const [email, setEmail] = useState<string | undefined>();
    const [password, setPassword] = useState<string | undefined>();

    const loginSubmit = async (email:string|undefined, password:string|undefined) => {
        setLoading(true);
        const isLoggedIn:boolean = await login(email, password)

        setLoggedIn(isLoggedIn);
        setLoading(false);
    }


    useEffect(() => {
        const getData = async () => {
            const resp:Response = await pingServer();
            const ping:PingModel = new PingModel(await resp.json());
            setPing(ping);
            setLoading(false);
        }
        const authUser = async () => {
            auth().then((isLoggedIn):void => {
                if(isLoggedIn) {
                    setLoggedIn(true);
                } else {
                    setLoggedIn(false);
                }   
                
            });
        }

        authUser();
        getData();
    }, []);

    return (
        <div className={styles.container}>
            {loading! === false 
                ? (
                    <>
                        <span>{ping?.toJSON()}</span>
                        {!loggedIn
                            ? (
                                <LoginForm
                                    className="home-login-form"
                                    setEmail={setEmail}
                                    setPassword={setPassword}
                                    onSubmit={() => loginSubmit(email, password)}
                                />
                            ):(
                                <></>
                            )
                        }
                    </>
                ):(
                    <span>Loading...</span>
                )
            }
        </div>
        
    );
};
