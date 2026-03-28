<?php
namespace App\Entity;
use Doctrine\ORM\Mapping as ORM;

#[ORM\Entity]
#[ORM\Table(name: 'orders')]
class Order {
    #[ORM\Id]
    #[ORM\GeneratedValue]
    #[ORM\Column(type: 'integer')]
    private int $id;

    #[ORM\Column(type: 'string', length: 255)]
    private string $customer_name;

    #[ORM\Column(type: 'decimal', precision: 10, scale: 2)]
    private string $total_amount;

    #[ORM\Column(type: 'string', length: 50)]
    private string $status;

    public function getId(): int { return $this->id; }
    public function getCustomerName(): string { return $this->customer_name; }
    public function setCustomerName(string $customer_name): self { $this->customer_name = $customer_name; return $this; }
    public function getTotalAmount(): string { return $this->total_amount; }
    public function setTotalAmount(string $total_amount): self { $this->total_amount = $total_amount; return $this; }
    public function getStatus(): string { return $this->status; }
    public function setStatus(string $status): self { $this->status = $status; return $this; }
}
