/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-17 23:39:43
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-19 21:34:46
 * @FilePath: /WindLiberity/Wind-Liberity/src/App.tsx
 * @Description: 
 * 
 */
import { Layout } from "antd"
import { Routes, Route } from "react-router-dom"
import NavBar from "./components/NavBar"
import Home from "./pages/Home"
import Login from "./pages/Login"
import Register from "./pages/Register"
import WeekSchedules from "./components/WeekSchedules";

const { Content } = Layout;

function App() {
  return (
    <Layout>
      <NavBar />
      <Content>
        <Routes>
          <Route
            path="/"
            element={
              <>
                <Home />
                <WeekSchedules />
              </>
            }
          />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
        </Routes>
      </Content>
    </Layout>
  )
}

export default App
