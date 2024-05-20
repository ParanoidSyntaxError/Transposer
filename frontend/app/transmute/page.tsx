import AppBar from "@/components/app-bar";
import Image from "next/image";
import TransmuteCard from "@/components/transmute-card";

export default function TransmutePage() {
    return (
        <div
            className="w-full p-4"
        >
            <div
                className="w-full"
            >
                <AppBar/>
            </div>
            <div
                className="flex flex-row items-center w-fit mx-auto mt-12 mb-4"
            >
                <div
                    className="text-4xl font-semibold"
                >
                    transmute
                </div>
                <Image
                    src="/PortalGreen.png"
                    alt=""
                    width={64}
                    height={64}
                />
            </div>
            <div
                className="max-w-96 w-full mx-auto"
            >
                <TransmuteCard/>
            </div>
        </div>
    );
}