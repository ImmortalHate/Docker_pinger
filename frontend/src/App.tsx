import React, { useEffect, useState } from "react";
import axios from "axios";
import { format } from "date-fns";
import { ru } from "date-fns/locale";
import "./App.css";

interface ContainerStatus {
  id: number;
  ip_address: string;
  ping_time: number;
  last_success_attempt: string;
}

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

const formatDate = (isoString: string): string => {
  try {
    return format(new Date(isoString), "dd.MM.yyyy HH:mm:ss", { locale: ru });
  } catch (error) {
    return isoString;
  }
};

const App: React.FC = () => {
  const [data, setData] = useState<ContainerStatus[]>([]);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API_URL}/api/status`);
      setData(response.data);
    } catch (error) {
      console.error("Ошибка загрузки данных:", error);
    }
  };

  useEffect(() => {
    fetchData();
    const intervalId = setInterval(fetchData, 5000);
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h1>Мониторинг контейнеров</h1>
      <table style={{ width: "100%", borderCollapse: "collapse" }}>
        <thead>
          <tr>
            <th style={{ border: "1px solid #ccc", padding: "8px" }}>ID</th>
            <th style={{ border: "1px solid #ccc", padding: "8px" }}>IP-адрес</th>
            <th style={{ border: "1px solid #ccc", padding: "8px" }}>Время пинга (мс)</th>
            <th style={{ border: "1px solid #ccc", padding: "8px" }}>Последняя проверка</th>
          </tr>
        </thead>
        <tbody>
          {data.map((status) => (
            <tr key={status.id}>
              <td style={{ border: "1px solid #ccc", padding: "8px" }}>{status.id}</td>
              <td style={{ border: "1px solid #ccc", padding: "8px" }}>{status.ip_address}</td>
              <td style={{ border: "1px solid #ccc", padding: "8px" }}>{status.ping_time}</td>
              <td style={{ border: "1px solid #ccc", padding: "8px" }}>
                {formatDate(status.last_success_attempt)}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default App;