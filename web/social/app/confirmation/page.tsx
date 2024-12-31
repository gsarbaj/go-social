'use client'

import {Button} from "@/components/ui/button";
import {handleConfirm} from "@/app/server/handleConfirm";
import {toast} from "sonner";
import {useRouter} from "next/navigation";



export default function ConfirmationPage () {


    const router = useRouter()


    const handleClick = () => {
        handleConfirm().then(r => {

            if (r.error) {
                toast.error(r.error, {
                    onAutoClose: () => {
                        console.log(r.error)
                    },
                    onDismiss: () => {
                        console.log(r.error)
                    }
                });
            } else {
                toast.success(r.message, {
                    onAutoClose: () => {
                        router.push("/");
                    },
                    onDismiss: () => {
                        router.push("/");
                    }
                });
            }
        });

    }

    return (
        <>
            <h1>Confirmation</h1>
            <Button onClick={handleClick}>Click to confirm</Button>
        </>
    );
};

// created on 31/12/2024 15:53
