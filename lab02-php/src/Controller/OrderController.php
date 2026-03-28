<?php
namespace App\Controller;

use App\Entity\Order;
use Doctrine\ORM\EntityManagerInterface;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;

class OrderController extends AbstractController {
    #[Route('/orders', methods: ['GET'])]
    public function index(EntityManagerInterface $em): JsonResponse {
        $orders = $em->getRepository(Order::class)->findAll();
        $data = array_map(fn($o) => ['id' => $o->getId(), 'customer_name' => $o->getCustomerName(), 'total_amount' => $o->getTotalAmount(), 'status' => $o->getStatus()], $orders);
        return new JsonResponse($data);
    }

    #[Route('/orders', methods: ['POST'])]
    public function create(Request $request, EntityManagerInterface $em): JsonResponse {
        $data = json_decode($request->getContent(), true);
        $order = new Order();
        $order->setCustomerName($data['customer_name']);
        $order->setTotalAmount($data['total_amount']);
        $order->setStatus($data['status']);
        $em->persist($order);
        $em->flush();
        return new JsonResponse(['id' => $order->getId()], 201);
    }

    #[Route('/orders/{id}', methods: ['GET'])]
    public function show(Order $order): JsonResponse {
        return new JsonResponse(['id' => $order->getId(), 'customer_name' => $order->getCustomerName(), 'total_amount' => $order->getTotalAmount(), 'status' => $order->getStatus()]);
    }

    #[Route('/orders/{id}', methods: ['PUT'])]
    public function update(Request $request, Order $order, EntityManagerInterface $em): JsonResponse {
        $data = json_decode($request->getContent(), true);
        $order->setCustomerName($data['customer_name']);
        $order->setTotalAmount($data['total_amount']);
        $order->setStatus($data['status']);
        $em->flush();
        return new JsonResponse();
    }

    #[Route('/orders/{id}', methods: ['DELETE'])]
    public function delete(Order $order, EntityManagerInterface $em): JsonResponse {
        $em->remove($order);
        $em->flush();
        return new JsonResponse();
    }
}
