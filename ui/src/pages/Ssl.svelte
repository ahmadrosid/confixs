<script lang="ts">
    import Button from "../components/ui/Button.svelte";
    import RetroOutline from "../components/ui/RetroOutline.svelte";
    import PlusIcon from 'lucide-svelte/icons/plus';
    import BaseLayout from "../components/BaseLayout.svelte";
    
    // Sample data for tasks
    let tasks = [
        { id: 1, title: "Review and merge pull requests", 
        status: "In Progress" },
        { id: 2, title: "Update user authentication module", 
        status: "Pending" },
        { id: 3, title: "Conduct code review for Project-B", 
        status: "Completed" },
        { id: 4, title: "Optimize database queries", 
        status: "Pending" }
    ];

    // Function to determine status color
    function getStatusColor(status: string) {
        switch (status.toLowerCase()) {
            case 'completed':
                return 'text-green-600';
            case 'in progress':
                return 'text-blue-600';
            case 'pending':
                return 'text-yellow-600';
            default:
                return 'text-gray-600';
        }
    }
</script>

<style>
    .tasks-table {
        width: 100%;
        border-collapse: collapse;
    }
    .tasks-table th, .tasks-table td {
        padding: 12px;
        border: none;
        text-align: left;
    }
</style>

<BaseLayout>
    <div class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"></div>

    <div class="container mx-auto max-w-6xl px-4 pb-12 relative">
        <div class="py-12 pb-4 flex justify-between items-center">
            <h1 class="text-4xl font-bold">SSL</h1>
            <Button> <PlusIcon class="mr-1 w-4 h-4"/> Add SSL</Button>
        </div>

        <div class="py-8">
            <RetroOutline className="w-full" childClassName="p-0">
                <div class="">
                    <table class="tasks-table table-auto">
                        <thead class="border-b border-gray-800">
                            <tr>
                                <th>ID</th>
                                <th>Title</th>
                                <th>Status</th>
                                <th class="w-64">Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#each tasks as task}
                                <tr>
                                    <td>{task.id}</td>
                                    <td>{task.title}</td>
                                    <td class={getStatusColor(task.status)}>{task.status}</td>
                                    <td>
                                        <Button className="mr-2" variant="outline-primary">View</Button>
                                        <Button className="mr-2" variant="warning">Edit</Button>
                                        <Button variant="danger">Delete</Button>
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                    </table>
                </div>
            </RetroOutline>
        </div>
    </div>
</BaseLayout>