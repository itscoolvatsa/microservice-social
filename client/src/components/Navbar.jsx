import React from "react";
import style from "./Navbar.module.scss";
import {
    AllInclusive,
    Search,
    Notifications,
    Chat,
    Settings,
    AccountCircle,
} from "@mui/icons-material";

const Navbar = () => {
    return (
        <header className={style.header}>
            <div className={style.left}>
                <AllInclusive className={style.icon} />
                <span>Micro Social</span>
            </div>
            <div className={style.middle}>
                <Search />
                <input type="text" />
            </div>
            <div className={style.right}>
                <Notifications />
                <Chat />
                <Settings />
                <AccountCircle />
            </div>
        </header>
    );
};

export default Navbar;
