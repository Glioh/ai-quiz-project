import { GameState, NetService, PacketTypes, type ChangeGameStatePacket, type Packet } from "../net";
import { writable, type Writable } from "svelte/store";

export const state: Writable<GameState> = writable(GameState.Lobby);

export class PlayerGame {
    private net: NetService;

    constructor() {
        this.net = new NetService();
        this.net.connect();
        this.net.onPacket(p => this.onPacket(p));
    }

    join(code: string, name: string) {
        let packet = {
            id: PacketTypes.Connect,
            code: code,
            name: name,
        };

        this.net.sendPacket(packet);
    }

    onPacket(packet: Packet) {
        switch(packet.id) {
            case PacketTypes.ChangeGameState: {
                let data = packet as ChangeGameStatePacket;
                state.set(data.state);
                break;
            }
            // Handle other packet types as needed
        }
    }
}
