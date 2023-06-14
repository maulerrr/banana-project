import React, {useEffect, useState} from 'react';
import LogOut from "../handlers/auth/HandleLogOut";
import classes from "./classes/components.module.css"
import {faBars, faRightFromBracket, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Cookies, useCookies} from "react-cookie";

function Navbar(props) {
    const [clicked, setClicked] = useState(false)
    const [dropdown, setDropdown] = useState(<></>)
    const [cookies, setCookie, removeCookie] = useCookies(["token"])

    const [showBurger, setShowBurger] = useState(false)
    const [burger, setBurger] = useState(<></>)
    const [menu, setMenu] = useState(<></>)

    const [windowSize, setWindowSize] = useState([
        window.innerWidth,
        window.innerHeight,
    ]);

    useEffect(()=> {
        window.onresize = (e) => setWindowSize([window.innerWidth, window.innerHeight])

        if (windowSize[0] < 768) {
            setShowBurger(true)
            setBurger(
                <div className={classes.BurgerContainer}>
                    <button onClick={handleShowBurger}>
                        <FontAwesomeIcon icon={faBars} color={"white"}/>
                    </button>
                    {menu}
                </div>
            )
        } else {
            setShowBurger(false)
            setBurger(<></>)
            setMenu(<></>)
        }
    }, [windowSize, clicked])

    function handleShowBurger(e) {
        if (clicked) setClicked(false)
        else setClicked(true)

        if (clicked) {
            setMenu(
                <ul>
                    <li>
                        <h1>Banana App 🍌</h1>
                    </li>
                    <li>
                        <a href="/">Home</a>
                    </li>
                    <li>
                        <a href="/create">Create</a>
                    </li>
                    <li>
                        { username ?
                            <button onClick={handleLogOut}>
                                Log Out
                                &nbsp;
                                <FontAwesomeIcon icon={faRightFromBracket}/>
                            </button>
                            : <></>}

                    </li>
                </ul>
            )
        } else setMenu(<></>)
    }

    function showDropDown(e){
        if (clicked) setClicked(false)
        else setClicked(true)

        if (username) {
            if (clicked) {
                setDropdown(
                    <button className={classes.LogOutButton} onClick={handleLogOut}>
                        Log out
                        &nbsp;
                        <FontAwesomeIcon icon={faRightFromBracket}/>
                    </button>
                )
            } else setDropdown(<></>)
        } else window.location.href = "/login"
    }

    function handleLogOut(e){
        setClicked(false)
        removeCookie("token")

        LogOut().then(()=>{
            console.log("see ya")
        }).finally(()=>{
            window.location.href='/login'
        })
    }

    const username = JSON.parse(localStorage.getItem("user"))?.username

    return (
        <header>
            {showBurger ? burger :
                <nav className={classes.Navbar}>
                    <h3>
                        Banana 🍌
                    </h3>

                    <a href="/">Home</a>

                    <a href="/create">Create</a>

                    <div onClick={showDropDown}>
                        <p className={classes.NavDropdown} >
                            <FontAwesomeIcon icon={faUser}/>
                            &nbsp;
                            {
                                username ? username : <a href={"/login"}>Login</a>
                            }
                        </p>
                        {dropdown}
                    </div>
                </nav>
            }

        </header>
    );
}

export default Navbar;