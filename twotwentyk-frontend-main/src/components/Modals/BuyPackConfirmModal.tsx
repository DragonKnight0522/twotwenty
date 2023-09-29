import React from "react";
import { Modal as ModalWrapper } from "./Modal";
import { SellModalProps } from "../../types";
import { ButtonGroup, BuyPackConfirmModalWrapper, IconWrapper } from "./styles";
import { Button } from "../Button";
import { useNavigate } from "react-router-dom";
import { IconConfirm } from "../Icons";

export const BuyPackConfirmModal: React.FC<SellModalProps> = ({
  onClose,
  open,
  isMarket,
}) => {
  const navigate = useNavigate();
  return (
    <ModalWrapper onClose={onClose} open={open} width={407}>
      <BuyPackConfirmModalWrapper>
        {isMarket ? (
          <IconWrapper>
            <img src="/assets/buy-success.png" alt="" />
          </IconWrapper>
        ) : (
          <h3>Success!</h3>
        )}
        <h3>Congratulations!</h3>
        <p>
          Your {isMarket ? "purchase" : "transaction"} was completed
          successfully.
        </p>
        <ButtonGroup>
          <Button
            onClick={() =>
              isMarket ? onClose() : navigate("/dashboard/packs")
            }
          >
            {isMarket ? "Buy More" : "Buy More Packs"}
          </Button>
          <Button
            variant="outlined"
            onClick={() => navigate("/dashboard/dates")}
          >
            {isMarket ? "View Purchase" : "  Open Packs"}
          </Button>
        </ButtonGroup>
      </BuyPackConfirmModalWrapper>
    </ModalWrapper>
  );
};
