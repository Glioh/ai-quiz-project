<script lang="ts">
    import type { PlayerGame } from "../../service/player/player";

    export let game: PlayerGame;
</script>

<div class="h-screen bg-purple-500 relative">
    <!-- Connected Status with Icon -->
    <div class="absolute top-4 left-4 flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
        <span class="text-white">Connected</span>
    </div>

    <!-- Main Container -->
    <div class="flex items-center justify-center h-full">
        <div class="max-w-md w-full bg-white/10 backdrop-blur-sm rounded-lg p-6 mx-4">

            <!-- Main Content -->
            <div class="text-center mb-4">
                <h1 class="text-2xl font-bold text-white mb-2">Welcome to the Game!</h1>
                <p class="text-white/80">Waiting on Host to Start...</p>
            </div>

            <!-- Loading Animation -->
            <div class="flex justify-center gap-2">
                <div class="w-2 h-2 bg-white rounded-full animate-pulse"></div>
                <div class="w-2 h-2 bg-white rounded-full animate-pulse"></div>
                <div class="w-2 h-2 bg-white rounded-full animate-pulse"></div>
            </div>
        </div>
    </div>
</div>