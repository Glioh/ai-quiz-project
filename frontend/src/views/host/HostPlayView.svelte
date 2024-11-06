<script lang="ts">
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import Clock from "../../lib/Clock.svelte";
    import { COLORS, type QuizChoice } from "../../model/quiz";
    import AudioController from "../../lib/AudioController.svelte";
    import { type HostGame, tick, currentQuestion, state } from "../../service/host/host";
    import { GameState } from "../../service/net";

    function getCardColor(choice: QuizChoice, state: GameState, defaultColor: string){
        if(state != GameState.Reveal)
            return defaultColor
        return choice.correct ? "bg-green-400" : "bg-red-400"
    }
    
    function skipQuestion() {
        game.skip();
    }

    export let game: HostGame;

    // Determine which audio file to play based on game state
    $: audioFile = $state === GameState.Reveal ? "Reveal.mp3" : "Play.mp3";
</script>

{#if $currentQuestion != null}
    <div class="min-h-screen h-screen flex flex-col">
        <div class="absolute left-4 top-4 z-10">
            <AudioController 
                {audioFile}
                iconColor="black"
                hoverBgColor="bg-gray-100"
            />
        </div>

        <div class="absolute right-4 top-4 z-10">
            <button 
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                on:click={skipQuestion}
            >
                Skip
            </button>
        </div>

        <div class="bg-white text-3xl border-b p-4 font-bold text-center">
            {$currentQuestion.name}
        </div>
        <div class="flex-1 flex flex-col justify-center pl-4">
            <div class="flex justify-between items-center">
                <Clock>
                    <span class="text-3xl">{$tick}</span>
                </Clock>
                <img
					src="https://placehold.co/500x250"
					alt="center"
					class="max-w-[500px]"
				/>
                <div class="w-24"></div>
            </div>
        </div>
        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard color={getCardColor($currentQuestion.choices[i], $state, color)}>
                    <p class="pl-14">{$currentQuestion.choices[i].name}</p>
                </QuizChoiceCard>
            {/each}
        </div>
    </div>
{/if}