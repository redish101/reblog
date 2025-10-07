"use client";

import { env } from "@/lib/env";
import { removeMidAddress } from "@/lib/utils";
import { ShieldCheck } from "lucide-react";
import Link from "next/link";

interface CopyrightCardProps {
  blockNumber: number;
  owner: string;
  createTxHash: string;
  updateTxHash?: string;
  ipfsURL: string;
}

export function CopyrightCard(props: CopyrightCardProps) {
  const infoItems = [
    {
      label: "区块编号",
      value: <div>#{props.blockNumber}</div>,
    },
    {
      label: "所有者",
      value: (
        <Link
          href={`${env.blockExplorerBaseURL}/address/${props.owner}`}
          target="_blank"
        >
          {props.owner}
        </Link>
      ),
    },
    {
      label: "交易哈希",
      value: (
        <div>
          <Link
            href={`${env.blockExplorerBaseURL}/tx/${props.createTxHash}`}
            target="_blank"
            className="inline-block mr-4 break-all focus-visible:ring focus-visible:ring-accent"
          >
            发布&nbsp;{removeMidAddress(props.createTxHash)}
          </Link>
          {props.updateTxHash && (
            <Link
              href={`${env.blockExplorerBaseURL}/tx/${props.updateTxHash}`}
              target="_blank"
              className="inline-block break-all focus-visible:ring focus-visible:ring-accent"
            >
              最后更新&nbsp;{removeMidAddress(props.updateTxHash)}
            </Link>
          )}
        </div>
      ),
    },
    {
      label: "IPFS 地址",
      value: (
        <Link
          href={(() => {
            const cid = props.ipfsURL.replace("ipfs://", "");
            return `https://ipfs.io/ipfs/${cid}`;
          })()}
        >
          {props.ipfsURL}
        </Link>
      ),
    },
  ];

  return (
    <div className="text-sm bg-zinc-50 rounded-xl py-6 p-8">
      <span className="flex items-center text-gray-900 mb-4">
        <ShieldCheck className="mr-1 text-green-600 inline-flex items-center w-5 h-5" />
        此文章数据所有权由区块链加密技术和智能合约保障仅归创作者所有。
      </span>
      <ul className="space-y-2 text-[13px] text-gray-500 overflow-hidden">
        {infoItems.map((item) => (
          <li key={item.label}>
            <div className="font-medium">{item.label}</div>
            {item.value}
          </li>
        ))}
      </ul>
    </div>
  );
}
