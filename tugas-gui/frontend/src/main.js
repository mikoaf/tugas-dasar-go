import './style.css';
import './app.css';

import logo from './assets/images/nba.png';
// import {Greet} from '../wailsjs/go/main/App';

import { GetSensorData } from '../wailsjs/go/api/Api'

document.querySelector('#app').innerHTML = `
    <img id="logo" class="logo">
    <h1>Sensor Info</h1>
    <div>
        <pre id="output">Memuat...</pre>
        <button id="refresh">Refresh</button>
    </div>
`;
document.getElementById('logo').src = logo;

async function UpdateSensorData() {
    try{
        const data = await GetSensorData();
        document.getElementById('output').innerText = data;
    } catch (err) {
        console.error("failed to get sensor data:", err);
        document.getElementById('output').innerText = "failed to get sensor data.";
    }
}

document.getElementById('refresh').addEventListener('click', UpdateSensorData);

UpdateSensorData();

setInterval(UpdateSensorData, 1000);
