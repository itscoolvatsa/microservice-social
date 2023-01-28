import React, {useState} from "react";
import style from "./login.module.scss";
import {AllInclusive} from "@mui/icons-material";
import {signIn} from "next-auth/react";

const Login = () => {
    const [inputs, setInputs] = useState({
        email: "",
        password: "",
    });
    const [err, setErr] = useState(null);

    const handleChange = (e) => {
        setInputs((prev) => ({...prev, [e.target.name]: e.target.value}));
    };

    const handleLogin = async (e) => {
        const status = await signIn("credentials", {
            redirect: false,
            email: inputs.email,
            password: inputs.password,
            callbackUrl: "/",
        });

        if (status.ok) router.push(status.url);
    };

    return (
        <main className={style.login}>
            <div className={style.main}>
                <div className={style.header}>
                    <AllInclusive/>
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
