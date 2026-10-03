/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-17 23:58:59
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-18 19:50:43
 * @FilePath: /app/WindLiberity/Wind-Liberity/src/pages/Home.tsx
 * @Description: 
 * 
 */
import { Carousel } from "antd";
import React from "react";
import Reze from "../assets/Reze.jpeg"
const banners = [
  { src: Reze, alt: "banner-1" },
  { src: Reze, alt: "banner-2" },
  { src: Reze, alt: "banner-3" },
  { src: Reze, alt: "banner-4" },
]

const animeList_weekly = 
[
    { title: ""}
]

const Home = () => (
    <Carousel autoplay>
        {banners.map((item) =>(
            <img
            src={item.src}
            alt={item.alt}
            style={{ width: "100%", height: "100%", objectFit: "cover"}}
        />
        ))}
    </Carousel>
);

export default Home;