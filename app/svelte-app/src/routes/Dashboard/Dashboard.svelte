<script lang="ts">
    import "./style.css"
    import { Chart, Line, Bubble, Pie } from 'svelte-chartjs'
    import {
    Chart as ChartJS,
    Title,
    Tooltip,
    Legend,
    PointElement,
    LinearScale,
    ArcElement,
    CategoryScale,
    LineElement
  } from 'chart.js';
  ChartJS.register(Title, Tooltip, Legend, LineElement, LinearScale, PointElement, CategoryScale);


  ChartJS.register(Title, Tooltip, Legend, ArcElement, CategoryScale)
  ChartJS.register(Title, Tooltip, Legend, PointElement, LinearScale);
    const data = {
  labels: 'Bubble',
  datasets: [
    {
      label: 'John',
      data: [
        {
          x: 3,
          y: 7,
          r: 10,
        },
      ],
      backgroundColor: '#ff6384',
      hoverBackgroundColor: '#ff6384',
    },
    {
      label: 'Peter',
      data: [
        {
          x: 3.2,
          y: 7,
          r: 10,
        },
      ],
      backgroundColor: '#44e4ee',
      hoverBackgroundColor: '#44e4ee',
    },
    {
      label: 'Donald',
      data: [
        {
          x: 3.4,
          y: 7,
          r: 10,
        },
      ],
      backgroundColor: '#62088A',
      hoverBackgroundColor: '#62088A',
    },
  ],
};
</script>

<main>
  <section>
    <h1>Greetings</h1>
    <p>Hey there! Welcome to the dashboard, you can view your data here.</p>
  </section>
  <section class="data">
    <h1>Graph</h1>
 
  </section>
  <section class="data">
    <h1>Line</h1>
 
  </section>
  <section class="data">
    <h1>Pie</h1>
 
  </section>
  <section class="data">
    <h1>Donut</h1>
  
  </section>
  <div class="spacer">
    <hr>
    <p>@denzeysenpai</p>
  </div>
</main> 
