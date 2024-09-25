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
  </section>
  <!-- <section class="data">
    <h1>Data Section</h1>
    <Pie
    {data}
    height={80}
    options={{ maintainAspectRatio: true }}
    />   
  </section> -->
</main>
