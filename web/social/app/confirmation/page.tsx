'use client'

import {Button} from "@/components/ui/button";
import {handleConfirm} from "@/app/server/handleConfirm";
import {useSearchParams} from "next/navigation";

export default function ConfirmationPage () {


    const handleClick = () => {
        handleConfirm().then(r => console.log(r));
    }

    return (
        <>
            <h1>Confirmation</h1>
            <Button onClick={handleClick}>Click to confirm</Button>
        </>
    );
};

// created on 31/12/2024 15:53
