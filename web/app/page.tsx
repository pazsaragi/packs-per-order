"use server"

import { OrderForm, OrderInput } from "@/components/order-form";
import Image from "next/image";
import { submitOrder } from '@/app/actions';

export default async function Home() {

  const handleOrder = async (values: OrderInput) => {
    "use server"
    return submitOrder(values)
  }

  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <OrderForm submitOrder={handleOrder} />
      </main>
      <footer className="row-start-3 flex gap-6 flex-wrap items-center justify-center">
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://packs-per-order.onrender.com"
          target="_blank"
          rel="noopener noreferrer"
        >
         API Url
        </a>
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://github.com/pazsaragi/packs-per-order/blob/main/backend/api.http"
          target="_blank"
          rel="noopener noreferrer"
        >
          <Image
            aria-hidden
            src="https://nextjs.org/icons/window.svg"
            alt="Window icon"
            width={16}
            height={16}
          />
          Examples
        </a>
        <a
          className="flex items-center gap-2 hover:underline hover:underline-offset-4"
          href="https://vercel.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          Deployed with vercel.com and render.com →
        </a>
      </footer>
    </div>
  );
}
