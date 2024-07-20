<script lang="ts">
  import { createQuery } from "@tanstack/svelte-query";
  import { Bar } from "svelte-chartjs";
  import {
    Chart,
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale,
  } from "chart.js";

  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import { pb } from "../lib/pocketbase";

  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";
  import Loading from "./ui/Loading.svelte";

  let name = "";
  Chart.register(
    Title,
    Tooltip,
    Legend,
    BarElement,
    CategoryScale,
    LinearScale
  );
  let chartData = {
    labels: ["CPU", "Memory", "Disk"],
    datasets: [
      {
        label: "Usage",
        data: [0, 0, 0], 
        backgroundColor: "#000",
      },
    ],
  };

  let chartOptions = {
    responsive: true,
    maintainAspectRatio: false,
  };

  onMount(async () => {
    if (!pb.authStore.isValid) {
      navigate("/login");
      return;
    }

    const user = pb.authStore.model;
    if (user) {
      name = user.name;
    }
  });

  type ResourceUsageResponse = {
    cpu: number;
    memory: number;
    disk: number;
  };

  const fetchResourceUsage = async (): Promise<ResourceUsageResponse> => {
    const response = await fetch("/api/resource/usage");
    if (!response.ok) {
      throw new Error("Failed to fetch NGINX configs");
    }
    const data = await response.json();
    chartData.datasets[0].data = [
      parseFloat(data.cpu.toFixed(2)),
      parseFloat(data.memory.toFixed(2)),
      parseFloat(data.disk.toFixed(2))
    ];

    return {
      cpu: data.cpu,
      memory: data.memory,
      disk: data.disk,
    };
  };

  const resourceUsageQuery = createQuery({
    queryKey: ["resourceUsage"],
    queryFn: fetchResourceUsage,
  });
</script>

<RetroOutline childClassName="p-0">
  <div class="text-xl font-bold border-b border-gray-800 p-3">
    Server Resource Usage
  </div>
    {#if $resourceUsageQuery.isLoading}
        <div class="p-4 flex justify-center">
            <Loading />
        </div>
    {:else if $resourceUsageQuery.isError}
        <div class="text-rose-500 p-4">Error: {$resourceUsageQuery.error.message}</div>
    {:else if $resourceUsageQuery.isSuccess}
        <div class="p-4">
            <Bar data={chartData} options={chartOptions} />
        </div>
    {/if}
</RetroOutline>
