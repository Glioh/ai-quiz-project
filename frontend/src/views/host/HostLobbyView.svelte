<script lang="ts">
    import Button from '../../lib/Button.svelte';
    import type {HostGame} from "../../service/host/host";

    export let game: HostGame;

    function startGame() {
        game.start();
    }
</script>

<div class="p-8">
    <h2 class="flex justify-end">Lobby</h2>
    <div class="mt-4">
        <Button on:click={startGame}>Start Game</Button>
    </div>
</div>