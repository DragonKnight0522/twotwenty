import React, { useEffect, useState } from "react";
import { Modal as ModalWrapper } from "./Modal";
import type { WithdrawModalProps } from "../../types";
import { BalanceForWithdrawModalWrapper, IconWrapper } from "./styles";
import { IconCoins, IconCoinsLarge, IconPigMoney } from "../Icons";
import { Button } from "../Button";
import { useMyInfoContext } from "../../context";

export const BalanceForWithdrawModal: React.FC<WithdrawModalProps> = ({
  onClose,
  onWithdraw,
  open,
}) => {
  const { myInfoContext } = useMyInfoContext();
  const [balance, setBalance] = useState(0);

  useEffect(() => {
    setBalance(myInfoContext?.balance ? myInfoContext?.balance / 100 : 0);
  }, [myInfoContext]);
  return (
    <ModalWrapper onClose={onClose} open={open} width={408}>
      <BalanceForWithdrawModalWrapper>
        <h3>My Balance</h3>
        <IconWrapper>
          <IconCoinsLarge />
        </IconWrapper>
        <p>Your account balance is: </p>
        <h4>
          $
          {balance.toLocaleString("en-US", {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
          })}
          <span>USD</span>
        </h4>
        <Button className="withdraw-button" onClick={() => onWithdraw()}>
          Withdraw
        </Button>
      </BalanceForWithdrawModalWrapper>
    </ModalWrapper>
  );
};
