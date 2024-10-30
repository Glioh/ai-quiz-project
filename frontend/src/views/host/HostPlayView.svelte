<script lang="ts">
   import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
   import { type HostGame, tick, currentQuestion, state } from "../../service/host/host";

    export const COLORS = ["bg-pink-400", "bg-blue-400", "bg-yellow-400", "bg-purple-400"];

    export let game: HostGame;
</script>

{#if $currentQuestion != null}
    <div class="min-h-screen h-screen flex flex-col">
        <div class="bg-white text-3xl border-b p-4 font-bold text-center">
            {$currentQuestion.name}
        </div>
        <div class="flex-1 flex flex-col justify-center pl-4">
            <div
                class="bg-purple-500 text-white text-3xl p-4 rounded-lg mb-4"
            >
                {$tick}
            </div>
            <img
                alt="center"
                class="max-w-[500px]"
            />
            <div class="w-24"></div>
        </div>
        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard {color} choice={$currentQuestion.choices[i]} />
            {/each}
        </div>
    </div>
{/if}