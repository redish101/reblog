import { CopyrightCard } from "@/components/copyright-card";
import { Button } from "@/components/ui/button";
import { api } from "@/lib/api";

export default async function Home() {
  const postList = await api.posts.postsList();
  const post = postList.data.data![0]!;
  return (
    <div>
      <CopyrightCard
        blockNumber={post.block_number!}
        owner={post.owner_address!}
        createTxHash={post.create_tx_hash!}
        updateTxHash={post.update_tx_hash!}
        ipfsURL={post.ipfs_url!}
      />
    </div>
  );
}
