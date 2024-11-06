<script lang="ts">
    import type { PlayerGame } from "../../service/player/player";

    export let game: PlayerGame;
</script>

<div class="min-h-screen bg-gradient-to-br from-indigo-600 via-purple-600 to-pink-500">
    <div class="container mx-auto px-4 py-8">
        <!-- Top Status Bar -->
        <div class="flex justify-between items-center mb-8">
            <div class="flex items-center space-x-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-green-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
                <span class="text-white font-medium">Connected</span>
            </div>
            <div class="flex items-center space-x-2">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
                <span class="text-white font-medium">Waiting for Players...</span>
            </div>
        </div>

        <!-- Main Content -->
        <div class="bg-white/10 backdrop-blur-lg rounded-xl p-8 shadow-xl">
            <div class="text-center mb-8">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 text-white mx-auto mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1H2a1 1 0 01-1-1v-3a1 1 0 011-1h1a2 2 0 100-4H2a1 1 0 01-1-1V7a1 1 0 011-1h3a1 1 0 001-1V4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1H2a1 1 0 01-1-1v-3a1 1 0 011-1h1a2 2 0 100-4H2a1 1 0 01-1-1V7a1 1 0 011-1h3a1 1 0 001-1V4z" />
                </svg>
                <h1 class="text-4xl font-bold text-white mb-2">Welcome to the Game!</h1>
                <p class="text-lg text-white/80">Looking for other players...</p>
            </div>

            <!-- Player Status -->
            <div class="bg-white/20 rounded-lg p-6 backdrop-blur-sm mb-8">
                <h2 class="text-xl font-semibold text-white mb-4">Your Player Card</h2>
                <div class="flex items-center space-x-4">
                    <div class="w-12 h-12 rounded-full bg-gradient-to-r from-pink-500 to-purple-500 flex items-center justify-center">
                        <span class="text-white font-bold text-lg">P</span>
                    </div>
                    <div>
                        <p class="text-white font-medium">Do you see your name?</p>
                        <p class="text-white/60 text-sm">Waiting for game to start...</p>
                    </div>
                </div>
            </div>

            <!-- Loading Animation -->
            <div class="flex justify-center">
                <div class="animate-pulse flex space-x-4">
                    <div class="w-3 h-3 bg-white rounded-full"></div>
                    <div class="w-3 h-3 bg-white rounded-full"></div>
                    <div class="w-3 h-3 bg-white rounded-full"></div>
                </div>
            </div>
        </div>
    </div>
</div>