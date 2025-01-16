import {
	Box,
	Dialog,
	DialogContent,
	DialogTitle,
	List,
	ListItem,
	ListItemText,
	Typography,
} from "@mui/material";
import { useState } from "react";
import useSWR from "swr";
import type { P, Scheduler } from "../types";

const fetcher = async (key: string) => {
	return await fetch(key).then((res) => res.json());
};

const Home = () => {
	const [open, setOpen] = useState(false);
	const [dialogContent, setDialogContent] = useState({
		title: "",
		content: {},
	});

	const handleClick = (title: string, content: object) => {
		setDialogContent({ title, content });
		setOpen(true);
	};

	const handleClose = () => {
		setOpen(false);
	};

	const {
		data: ps,
		error: psError,
		isLoading: psLoading,
	} = useSWR<P[]>("http://localhost:8080/p", fetcher, {
		refreshInterval: 3000,
	});
	const {
		data: scheduler,
		error: schedulerError,
		isLoading: schedulerLoading,
	} = useSWR<Scheduler>("http://localhost:8080/sched", fetcher, {
		refreshInterval: 3000,
	});

	if (psLoading || schedulerLoading)
		return <div className={"loading"}>Loading...</div>;
	if (psError || schedulerError)
		return <div className={"error"}>Error fetching data.</div>;
	if (!ps || !scheduler) return <div className={"error"}>No data found.</div>;

	return (
		<div>
			<Box
				sx={{
					display: "flex",
					justifyContent: "center",
					alignItems: "center",
					margin: 2,
				}}
			>
				<Typography variant="h4">schedt</Typography>
			</Box>
			<Box
				sx={{
					display: "flex",
					justifyContent: "center",
					alignItems: "center",
					margin: 2,
				}}
			>
				<Box
					sx={{
						display: "flex",
						flexDirection: "column",
						alignItems: "center",
						margin: 2,
					}}
				>
					<Typography variant="h6">Run Queue</Typography>
					<Box sx={{ display: "flex", flexWrap: "wrap" }}>
						{scheduler.runq?.map((g, index) => (
							<Box
								key={index}
								sx={{
									display: "flex",
									justifyContent: "center",
									alignItems: "center",
									width: 50,
									height: 50,
									borderRadius: "50%",
									backgroundColor: "lightgreen",
									margin: 1,
									cursor: "pointer",
								}}
								onClick={() =>
									handleClick(`G${g.goid} Details`, {
										waitreason: g.waitreason,
										annotations: g.annotations,
										status: g.status,
									})
								}
							>
								<Typography>{g.goid}</Typography>
							</Box>
						))}
					</Box>
				</Box>

				<Box
					sx={{
						display: "flex",
						flexDirection: "column",
						alignItems: "center",
						margin: 2,
					}}
				>
					<Typography variant="h6">Stack</Typography>
					<Box sx={{ display: "flex", flexWrap: "wrap" }}>
						{scheduler.stack?.map((g, index) => (
							<Box
								key={index}
								sx={{
									display: "flex",
									justifyContent: "center",
									alignItems: "center",
									width: 50,
									height: 50,
									borderRadius: "50%",
									backgroundColor: "lightcoral",
									margin: 1,
									cursor: "pointer",
								}}
								onClick={() =>
									handleClick(`G${g.goid} Details`, {
										waitreason: g.waitreason,
										annotations: g.annotations,
										status: g.status,
									})
								}
							>
								<Typography>{g.goid}</Typography>
							</Box>
						))}
					</Box>
				</Box>

				<Box
					sx={{
						display: "flex",
						flexDirection: "column",
						alignItems: "center",
						margin: 2,
					}}
				>
					<Typography variant="h6">No Stack</Typography>
					<Box sx={{ display: "flex", flexWrap: "wrap" }}>
						{scheduler.noStack?.map((g, index) => (
							<Box
								key={index}
								sx={{
									display: "flex",
									justifyContent: "center",
									alignItems: "center",
									width: 50,
									height: 50,
									borderRadius: "50%",
									backgroundColor: "lightcoral",
									margin: 1,
									cursor: "pointer",
								}}
								onClick={() =>
									handleClick(`G${g.goid} Details`, {
										waitreason: g.waitreason,
										annotations: g.annotations,
										status: g.status,
									})
								}
							>
								<Typography>{g.goid}</Typography>
							</Box>
						))}
					</Box>
				</Box>
			</Box>

			{ps.map((p, index) => (
				<Box
					key={index}
					sx={{
						display: "flex",
						flexDirection: "column",
						alignItems: "center",
						margin: 2,
					}}
				>
					<Typography variant="h6">P{p.id}</Typography>
					<Box
						sx={{
							display: "flex",
							justifyContent: "center",
							alignItems: "center",
							width: 100,
							height: 100,
							backgroundColor: "lightblue",
							cursor: "pointer",
						}}
						onClick={() =>
							handleClick("M Details", { procid: p.m.procid, id: p.m.id })
						}
					>
						<Typography>M</Typography>
					</Box>
					<Box
						sx={{
							display: "flex",
							justifyContent: "center",
							alignItems: "center",
						}}
					>
						<Typography>Local Run Queue</Typography>
					</Box>
					<Box sx={{ display: "flex", flexWrap: "wrap", marginTop: 2 }}>
						{p.runq.map((g, index) => (
							<Box
								key={index}
								sx={{
									display: "flex",
									justifyContent: "center",
									alignItems: "center",
									width: 50,
									height: 50,
									borderRadius: "50%",
									backgroundColor: "lightgreen",
									margin: 1,
									cursor: "pointer",
								}}
								onClick={() =>
									handleClick(`G${g.goid} Details`, {
										waitreason: g.waitreason,
										annotations: g.annotations,
										status: g.status,
									})
								}
							>
								<Typography>{g.goid}</Typography>
							</Box>
						))}
					</Box>
					<Dialog open={open} onClose={handleClose}>
						<DialogTitle>{dialogContent.title}</DialogTitle>
						<DialogContent>
							<List>
								{Object.entries(dialogContent.content).map(
									([key, value], index) => (
										<ListItem key={index}>
											<ListItemText primary={`${key}: ${value}`} />
										</ListItem>
									),
								)}
							</List>
						</DialogContent>
					</Dialog>
				</Box>
			))}
		</div>
	);
};

export default Home;
