/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-18 20:28:38
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-22 15:46:22
 * @FilePath: /WindLiberity/Web/src/components/WeekSchedules.tsx
 * @Description: 
 * 
 */
import { useState } from "react";
import styled from "@emotion/styled"

const days  = ["周一", "周二", "周三", "周四", "周五", "周六", "周日"];
const ACCENT = "#fb7299";
const Bar = styled.div`
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 12px 16px;
  background: #f7f7f7;
  border-radius: 8px;
`;

const Title = styled.span`
    font-weight: bold;
    front-size: 16px;
    white-space: nowrap;
`;


const DayTabs = styled.div`
    display: flex;
    align-items: center;
    gap: 8px;
`;

const DayTab = styled.button<{ active?: boolean}>`
    border: none;
    cursor: pointer;
    font-size: 14px;
    line-height: 1;
    padding: 8px 14px;
    border-redius: 999px;
    background: ${({active }) => (active ? ACCENT : "transparent")};
    color: ${({ active}) => (active ? "#fff" : "#000")};
    transition: background 0.2s ease, color 0.2s ease;
    &:hover {
        background: ${({ active }) => (active ? "#e8608b" : "#ececec")};
    }
`;

const WeekSchedules = () => {
    const [activeDay, setActiveDay] = useState(2);
    return (
        <Bar>
            <Title>追番周表</Title>
            <DayTabs>
                {days.map((day, index) => (
                    <DayTab
                    key={day}
                    active={index == activeDay}
                    onClick={() => setActiveDay(index)}
                    >
                        {day}

                    </DayTab>
                ))}
            
            </DayTabs>

            
        </Bar>

    )
};

export default WeekSchedules;





