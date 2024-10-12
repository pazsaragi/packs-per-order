'use server'

import { OrderInput } from '@/components/order-form'

interface PacksResponse {
    Packs: Map<number, number>
}

export interface OrderResponse {
    success?: boolean;
    data?: PacksResponse;
    error?: string;
} 

export async function submitOrder(values: OrderInput): Promise<OrderResponse> {
  const { order } = values;

  try {
    const response = await fetch(`${process.env.API_URL}/pack?order=${order}`, {
      method: 'GET'
    })

    if (!response.ok) {
      throw new Error('Network response was not ok')
    }

    const result = await response.json()
    return { success: true, data: result }
  } catch (error) {
    console.error('Error submitting order:', error)
    return { error: 'An error occurred while processing your order. Please try again.' }
  }
}
