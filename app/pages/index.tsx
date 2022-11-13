import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import {Version} from '../models/Version'
import {GetVersion} from './api/version';

export default function Home() {
    useEffect(() => {
        const getData = async () => {
            const result:Promise<any> = await GetVersion();
            if(result !== null) {
                console.log(new Version(await result));
            } else {
            }
        };
        getData();
    }, []);

    return (
        <div className={styles.container}>
      
        </div>
    );
};
