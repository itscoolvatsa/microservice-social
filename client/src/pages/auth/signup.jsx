import React, {useState} from "react";
import style from "./login.module.scss";
import {AllInclusive} from "@mui/icons-material";
import useRequest from "/hooks/use-request";

const Signup = () => {
    const [inputs, setInputs] = useState({
        name: "",
        email: "",
        password: "",
    });

    const [err, setErr] = useState(null);

    const handleChange = (e) => {
        setInputs((prev) => ({...prev, [e.target.name]: e.target.value}));
    };

    const handleLogin = async (e) => {
        e.preventDefault();
        await doRequest();
    };

    const {doRequest, errors} = useRequest({
        url: "http://127.0.0.1:8080/users/signup",
        method: "post",
        body: inputs,
        // onSuccess: () => Router.push("/"),
    });

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
                        placeholder="Name"
                        name="name"
                        onChange={handleChange}
                    />
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
                    {/* {err && err} */}
                    {errors}
                    <button onClick={handleLogin}>SignUp</button>
                </div>
            </div>
        </main>
    );
};

export default Signup;
