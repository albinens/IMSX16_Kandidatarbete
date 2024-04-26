import React from "react";
import { Card } from "@mui/material";

/**
 * PRIO
   * Average Room Occupancy Rate: Calculate the average percentage of occupancy for each room over a specific period (e.g., daily, weekly, monthly).
   * 
   * Peak Occupancy Time: Identify the time of day when each room reaches its maximum occupancy.
   * 
   * Average Occupancy Duration: Calculate the average length of time each room is occupied.
 * 
 * SECONDARY
  * Room Utilization Rate: Calculate the percentage of time each room is occupied over a given period (e.g., daily, weekly, monthly).
  * 
  * Busiest Rooms: Identify the rooms with the highest average occupancy rates.
  * 
  * Least Utilized Rooms: Identify the rooms with the lowest average occupancy rates.
  * 
  * Occupancy Patterns: Analyze occupancy patterns to identify recurring trends, such as peak hours or days of the week with high or low occupancy.
  * 
  * Comparison with Capacity: Compare the actual occupancy with the room's capacity to understand underutilized or overutilized rooms.
  * 
  * Trend Analysis: Analyze changes in room occupancy over time to identify long-term trends.
  * 
  * Predicted Occupancy Accuracy: Measure the accuracy of your occupancy predictions compared to actual occupancy data.
 */
const KpiBox = (props) => {
  return (
    <Card style={{
      padding: '16px',
      margin: '10px',
      textAlign: 'center',
      height: '160px',
    }}>
      <div className="kpi-box-header">
        <h2>{props.value} {props.unit}</h2>
      </div>
      <div className="kpi-box-content">
        <p>{props.title}</p>
      </div>
    </Card>
  );
}

export default KpiBox;