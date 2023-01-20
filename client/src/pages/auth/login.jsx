import React, { useState } from "react";
import style from "./login.module.scss";
import { AllInclusive } from "@mui/icons-material";
import Router from "next/router";

const Login = () => {
    const [inputs, setInputs] = useState({
        email: "",
        password: "",
    });
    const [err, setErr] = useState(null);

    const handleChange = (e) => {
        setInputs((prev) => ({ ...prev, [e.target.name]: e.target.value }));
    };

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            await login(inputs);
            Router.push("/");
        } catch (err) {
            setErr(err.response.data);
        }
    };

    return (
        <main className={style.login}>
            <div className={style.main}>
                <div className={style.header}>
                    <AllInclusive />
                    <h1>Micro Social</h1>
                    <p>Enjoy Being Social</p>
                </div>

                <div className={style.form}>
                    <input
                        type="text"
                        placeholder="Email"
                        name="email"
                        onChange={handleChange}
                    />
                    <input
                        type="password"
                        placeholder="Password"
                        name="password"
                        onChange={handleChange}
                    />
                    {err && err}
                    <button onClick={handleLogin}>Login</button>
                </div>
            </div>
        </main>
    );
};

export default Login;
