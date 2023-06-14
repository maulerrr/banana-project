import './App.css';
import React from "react";
import {BrowserRouter, Routes, Route} from "react-router-dom";
import CreatePostPage from "./pages/CreatePostPage";
import Homepage from "./pages/Homepage";
import LoginPage from "./pages/LoginPage";
import SignUpPage from "./pages/SignUpPage";
import {CookiesProvider} from "react-cookie";
import PostPage from "./pages/PostPage";

function App() {
  return (
      <CookiesProvider>
          <BrowserRouter>
              <Routes>
                  <Route path={"/"} element={<Homepage/>} />
                  <Route path={"/create"} element={<CreatePostPage/>} />
                  <Route path={"/login"} element={<LoginPage/>} />
                  <Route path={"/signup"} element={<SignUpPage/>} />
                  <Route path={"/post/:post_id"} element={<PostPage/>} />
              </Routes>
          </BrowserRouter>
      </CookiesProvider>

  );
}

export default App;
