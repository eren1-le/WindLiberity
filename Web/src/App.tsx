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
import NavBar from "./components/NavBar"
import Home from "./pages/Home"
import WeekSchedules from "./components/WeekSchedules";
// 根组件：目前只渲染首页

const { Content } = Layout;
function App() {
  return (
    <Layout>
      <NavBar />
      <Content>
        <Home />
      </Content>
      <Content>
        <WeekSchedules />
      </Content>
    </Layout>
)
}

export default App
