"use client"

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { useState } from 'react'
import { OrderResponse } from "@/app/actions"

interface props {
    submitOrder: (values: OrderInput) => Promise<OrderResponse>
}

export const orderSchema = z.object({
    order: z.coerce.number().int().positive()
})

export type OrderInput =  z.infer<typeof orderSchema>

export function OrderForm({ submitOrder }: props) {
    // 1. Define your form.
    const form = useForm
        <OrderInput>
        ({
            resolver: zodResolver(orderSchema),
            defaultValues: {
                order: 200
            }
        })
    const [response, setResponse] = useState<OrderResponse | null>(null)

    // 2. Define a submit handler.
    async function onSubmit(values: OrderInput) {
        const result = await submitOrder(values)
        setResponse(result)
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
                <FormField
                    control={form.control}
                    name="order"
                    render={({ field }) => (
                        <FormItem>
                            <FormLabel>Order</FormLabel>
                            <FormControl>
                                <Input {...field} type="number" />
                            </FormControl>
                            <FormDescription>
                                This is your order size.
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <Button type="submit">Submit</Button>
            </form>
            {response && (
                <div className="mt-4">
                    {response.success && response.data ? (
                        <div className="text-green-600">
                            <h3 className="font-semibold">Order processed successfully:</h3>
                            <ul className="list-disc list-inside">
                                {Object.entries(response.data.Packs)
                                .map(([packSize, quantity]: [string, number]) => (
                                    <li key={packSize}>{quantity} pack(s) of {packSize}</li>
                                ))}
                            </ul>
                        </div>
                    ) : (
                        <p className="text-red-600">{response.error}</p>
                    )}
                </div>
            )}
        </Form>
    )

}
