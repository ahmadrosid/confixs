<script>
    import RetroOutline from "./ui/RetroOutline.svelte";
    import { Check, AlertTriangle, XCircle, Clock } from "lucide-svelte";
  
    const sites = [
      { 
        name: "ahmadrosid.com", 
        status: "Online", 
        statusClass: "text-green-600",
        icon: Check,
        lastUpdate: new Date("2024-07-21T10:30:00")
      },
      { 
        name: "v1.ahmadrosid.com", 
        status: "Maintenance", 
        statusClass: "text-yellow-600",
        icon: AlertTriangle,
        lastUpdate: new Date("2024-07-21T09:15:00")
      },
      { 
        name: "dev.ahmadrosid.com", 
        status: "Offline", 
        statusClass: "text-red-600",
        icon: XCircle,
        lastUpdate: new Date("2024-07-21T08:45:00")
      }
    ];
  
    function formatLastUpdate(date) {
      return date.toLocaleString('en-US', { 
        month: 'short', 
        day: 'numeric', 
        hour: '2-digit', 
        minute: '2-digit' 
      });
    }
  </script>
  
  <RetroOutline childClassName="p-0">
    <header class="text-xl font-bold border-b border-gray-800 p-3">Site Status</header>
    <div class="p-4">
      <ul class="space-y-4">
        {#each sites as site}
          <li>
            <div class="font-semibold mb-2">{site.name}</div>
            <div class="flex gap-2 items-center text-xs">
              <div class="flex items-center gap-1 {site.statusClass}">
                <span>{site.status}</span>
              </div>
              <div class="flex items-center gap-1 text-gray-500">
                <span>Last Update:</span>
                <span>{formatLastUpdate(site.lastUpdate)}</span>
              </div>
            </div>
          </li>
        {/each}
      </ul>
    </div>
  </RetroOutline>