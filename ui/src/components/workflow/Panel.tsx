// components/AddNodePanel.tsx
import { ScrollArea } from "../ui/scroll-area";
import { Sheet, SheetContent, SheetTitle } from "../ui/sheet";

type AddNodePanelProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  children: React.ReactNode;
  name: string;
};

const Panel = ({ open, onOpenChange, children, name }: AddNodePanelProps) => {
  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className="sm:max-w-[640px] p-0">
        <SheetTitle className="bg-primary p-4 text-white">{name}</SheetTitle>

        <ScrollArea className="p-10 flex-col space-y-5 h-[90vh]">{children}</ScrollArea>
      </SheetContent>
    </Sheet>
  );
};

export default Panel;
